// Copyright (c) 2019 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package alias

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	yaml "gopkg.in/yaml.v2"

	"github.com/iotexproject/iotex-core/ioctl/config"
	"github.com/iotexproject/iotex-core/ioctl/output"
)

// Multi-language support
var (
	importCmdShorts = map[config.Language]string{
		config.English: "Import aliases",
		config.Chinese: "导入别名",
	}
	importCmdUses = map[config.Language]string{
		config.English: "import 'DATA'",
		config.Chinese: "import '数据'",
	}
	flagImportFormatUsages = map[config.Language]string{
		config.English: "set format: json/yaml",
		config.Chinese: "设置格式：json / yaml",
	}
	flagForceImportUsages = map[config.Language]string{
		config.English: "override existing aliases",
		config.Chinese: "覆盖现有别名",
	}
)

// aliasImportCmd represents the alias import command
var aliasImportCmd = &cobra.Command{
	Use:   config.TranslateInLang(importCmdUses, config.UILanguage),
	Short: config.TranslateInLang(importCmdShorts, config.UILanguage),
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true
		err := aliasImport(cmd, args)
		return output.PrintError(err)
	},
}

type importMessage struct {
	ImportedNumber int     `json:"importedNumber"`
	TotalNumber    int     `json:"totalNumber"`
	Imported       []alias `json:"imported"`
	Unimported     []alias `json:"unimported"`
}

func init() {
	aliasImportCmd.Flags().StringVarP(&format,
		"format=", "f", "json", config.TranslateInLang(flagImportFormatUsages, config.UILanguage))
	aliasImportCmd.Flags().BoolVarP(&forceImport,
		"force-import", "F", false, config.TranslateInLang(flagForceImportUsages, config.UILanguage))
}

func aliasImport(cmd *cobra.Command, args []string) error {
	var importedAliases aliases
	switch format {
	default:
		return output.NewError(output.FlagError, fmt.Sprintf("invalid format flag %s", format), nil)
	case "json":
		if err := json.Unmarshal([]byte(args[0]), &importedAliases); err != nil {
			return output.NewError(output.SerializationError, "failed to unmarshal imported aliases", err)
		}
	case "yaml":
		if err := yaml.Unmarshal([]byte(args[0]), &importedAliases); err != nil {
			return output.NewError(output.SerializationError, "failed to unmarshal imported aliases", err)
		}
	}
	aliases := GetAliasMap()
	message := importMessage{TotalNumber: len(importedAliases.Aliases), ImportedNumber: 0}
	for _, importedAlias := range importedAliases.Aliases {
		if !forceImport && config.ReadConfig.Aliases[importedAlias.Name] != "" {
			message.Unimported = append(message.Unimported, importedAlias)
			continue
		}
		for aliases[importedAlias.Address] != "" {
			delete(config.ReadConfig.Aliases, aliases[importedAlias.Address])
			aliases = GetAliasMap()
		}
		config.ReadConfig.Aliases[importedAlias.Name] = importedAlias.Address
		message.Imported = append(message.Imported, importedAlias)
		message.ImportedNumber++
	}
	out, err := yaml.Marshal(&config.ReadConfig)
	if err != nil {
		return output.NewError(output.SerializationError, "failed to marshal config", err)
	}
	if err := os.WriteFile(config.DefaultConfigFile, out, 0600); err != nil {
		return output.NewError(output.WriteFileError,
			fmt.Sprintf("failed to write to config file %s", config.DefaultConfigFile), err)
	}
	fmt.Println(message.String())
	return nil
}

func (m *importMessage) String() string {
	if output.Format == "" {
		line := fmt.Sprintf("%d/%d aliases imported\nExisted aliases:", m.ImportedNumber, m.TotalNumber)
		for _, alias := range m.Unimported {
			line += fmt.Sprint(" " + alias.Name)
		}
		return line
	}
	return output.FormatString(output.Result, m)
}
