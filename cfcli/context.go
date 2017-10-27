package cfcli

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"sync"
)

var context CFContext

type CFContext struct {
	cliConnection plugin.CliConnection
	cache         Cache
}

type Cache struct {
	sync.Mutex
	spaceList []string
	orgList   []string
}

func (c *Cache) Spaces() []string {
	c.Lock()
	defer c.Unlock()
	spaces := make([]string, len(c.spaceList), cap(c.spaceList))
	copy(spaces, c.spaceList)
	return c.spaceList
}

func (c *Cache) Orgs() []string {
	c.Lock()
	defer c.Unlock()
	orgs := make([]string, len(c.orgList), cap(c.orgList))
	copy(orgs, c.orgList)
	return orgs
}

func SetCFContext(cliConnection plugin.CliConnection) {
	context.cliConnection = cliConnection
	fmt.Println("Fetching visible Orgs...")
	orgs := listOrgs(cliConnection)
	fmt.Println("Fetching visible Spaces...")
	spaces := listSpaces(cliConnection)
	context.cache = Cache{spaceList: spaces, orgList: orgs}
}

func listSpaces(cliConnection plugin.CliConnection) []string {
	spaces, err := cliConnection.GetSpaces()
	if err != nil {
		return nil
	}
	spaceList := make([]string, 0, len(spaces))
	for _, space := range spaces {
		spaceList = append(spaceList, space.Name)
	}
	return spaceList
}

func listOrgs(cliConnection plugin.CliConnection) []string {
	orgs, err := cliConnection.GetOrgs()
	if err != nil {
		return nil
	}
	orgList := make([]string, 0, len(orgs))
	for _, org := range orgs {
		orgList = append(orgList, org.Name)
	}
	return orgList
}
