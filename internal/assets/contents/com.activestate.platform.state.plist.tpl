<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
    <dict>
        <key>Label</key>
        <string>my.everydaytasks</string>
        <key>ProgramArguments</key>
        <array>
            <!-- Wrap in `sh -c` so that $HOME is expanded -->
            <string>sh</string>
            <string>-c</string>
            {{- if .Args }}
            <string>{{.Exec}} {{.Args}}</string>
            {{- else}}
            <string>{{.Exec}}</string>
            {{- end}}
        </array>
        <key>ProcessType</key>
        <string>Interactive</string>
        <key>RunAtLoad</key>
        <true/>
        <key>KeepAlive</key>
        <false/>
    </dict>
</plist>