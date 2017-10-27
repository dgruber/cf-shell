package cfcli

import (
	"code.cloudfoundry.org/cli/command/common"
	"github.com/c-bata/go-prompt"
	"reflect"
	"strings"
)

func createCompletions() []prompt.Suggest {
	t := reflect.TypeOf(common.Commands)
	completions := make([]prompt.Suggest, 0, t.NumField()+len(extensions))
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Tag.Get("command") == "" {
			continue
		}
		completions = append(completions, prompt.Suggest{Text: f.Tag.Get("command"), Description: f.Tag.Get("description")})
	}
	for _, pluginCommand := range extensions {
		completions = append(completions, prompt.Suggest{Text: pluginCommand[0], Description: pluginCommand[1]})
	}
	return completions
}

func createUsageCompletionsMap() map[string]string {
	usageMap := make(map[string]string)
	t := reflect.TypeOf(common.Commands)
	for i := 1; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.Type.Kind() == reflect.Struct {
			usageVar, exists := f.Type.FieldByName("usage")
			if !exists {
				continue
			}
			usageMap[f.Tag.Get("command")] = usageVar.Tag.Get("usage")
		}
	}
	return usageMap
}

func createUsageCompletion(cmd string) []prompt.Suggest {
	var usageCompletion string
	usage := createUsageCompletionsMap()[cmd]
	for _, line := range strings.Split(usage, "\n") {
		if strings.Contains(line, "CF_NAME") {
			usageCompletion = strings.Replace(line, "CF_NAME", "", -1)
			usageCompletion = strings.Replace(usageCompletion, cmd, "", 1)
			usageCompletion = strings.TrimSpace(usageCompletion)
			break
		}
	}
	if StartsWithResolvableKeyWord(usageCompletion) {
		entities := ResolveKeyWord(usageCompletion)
		suggestions := []prompt.Suggest{{Text: " ", Description: usageCompletion}}
		for _, entity := range entities {
			suggestions = append(suggestions, prompt.Suggest{Text: entity})
		}
		return suggestions
	}
	return []prompt.Suggest{{Text: " ", Description: usageCompletion}}
}

func Completer(d prompt.Document) []prompt.Suggest {
	text := strings.Split(d.TextBeforeCursor(), " ")
	if len(text) > 1 {
		return createUsageCompletion(text[0])
	}
	return prompt.FilterContains(createCompletions(), text[0], true)
}
