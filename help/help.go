package help

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintHelp() {
	fmt.Println("OVERVIEW: the center linux container help interface.")
	color.Green("USAGE: ctcontainer-go <command> [<args>...]")
	fmt.Println()
	fmt.Println("OPTIONS:")
	fmt.Println("\tlist [installed [size]|available]\n\t\tList containers differentiated from [installed] and [available]")
	fmt.Println("\tload <path> <command>\n\t\tRun <command> in chroot, but you can set rootfs <path> custom")
	fmt.Println("\trun <codename> <command>\n\t\tRun <command> in the specified <codename> container")
	fmt.Println("\tinit <distribution>\n\t\tCreate the specified <distribution> container")
	fmt.Println("\tpurge <codename>\n\t\tDestroy the specified <codename> container")
	fmt.Println("\tversion|v\n\t\tShow the center container version")
	fmt.Println()
	fmt.Println("\t[args]\tOptional parameters for other settings")
	fmt.Println("\t\t-r/--ctroot=<path>\n\t\t-s/--ctshare=<path>\n\t\t-l/--ctlog_level=(debug|info)\n\t\t-u/--ctauto_umount=(0|1)\n\t\t-p/--ctsafety_level=(0|1|2)\n\t\t-b/--ctbackend=(chroot|nspawn)")
}
