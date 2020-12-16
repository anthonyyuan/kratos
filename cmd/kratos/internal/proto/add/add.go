package add

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/spf13/cobra"
	"golang.org/x/mod/modfile"
)

// CmdAdd represents the add command.
var CmdAdd = &cobra.Command{
	Use:   "add",
	Short: "Add a proto API template",
	Long:  "Add a proto API template. Example: kratos add helloworld/v1/hello.proto",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	// kratos add helloworld/v1/helloworld.proto
	input := args[0]
	n := strings.LastIndex(input, "/")
	path := input[:n]
	fileName := input[n+1:]
	pkgName := strings.ReplaceAll(path, "/", ".")

	p := &Proto{
		Name:        fileName,
		Path:        path,
		Package:     pkgName,
		GoPackage:   goPackage(pkgName),
		JavaPackage: javaPackage(pkgName),
		Service:     serviceName(fileName),
	}
	if err := p.Generate(); err != nil {
		fmt.Println(err)
		return
	}
}

func modName() string {
	modBytes, err := ioutil.ReadFile("go.mod")
	if err != nil {
		if modBytes, err = ioutil.ReadFile("../go.mod"); err != nil {
			return ""
		}
	}
	return modfile.ModulePath(modBytes) + "/api/"
}

func goPackage(name string) string {
	sub := strings.Split(name, ".")
	return modName() + strings.ReplaceAll(name, ".", "/") + ";" + sub[len(sub)-1]
}
func javaPackage(name string) string {
	return name
}

func serviceName(name string) string {
	return unexport(strings.Split(name, ".")[0])
}

func unexport(s string) string { return strings.ToUpper(s[:1]) + s[1:] }