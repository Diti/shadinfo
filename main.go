package main

import (
	"html/template"
	"net/http"

	"github.com/dustin/go-humanize"
)

var templates = template.Must(template.New("index.html.tmpl").Funcs(template.FuncMap{
	"prettyBytes": func(i uint64) string {
		return humanize.Bytes(i)
	},
}).ParseFiles("index.html.tmpl"))

func handler(w http.ResponseWriter, r *http.Request) {
	// Get system information with `gopsutil`
	cpuInfo, cpuErr := cpuInfo()
	diskInfo, diskErr := diskInfo()
	hostInfo, hostErr := hostInfo()
	memInfo, memErr := memInfo()
	// Abort on error with a HTTP 500 status code
	for _, err := range []interface{}{cpuErr, diskErr, hostErr, memErr} {
		if err != nil {
			http.Error(w, err.(error).Error(), http.StatusInternalServerError)
			return
		}
	}

	// Send system information to template
	info := make(map[string]interface{})
	info["Cpu"] = cpuInfo
	info["Disk"] = diskInfo
	info["Host"] = hostInfo
	info["RemoteAddr"] = r.RemoteAddr
	info["Mem"] = memInfo
	if err := templates.Execute(w, info); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8081", nil)
}
