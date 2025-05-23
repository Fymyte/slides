package code

// cmds: Multiple commands; placeholders can be used
// Placeholders <file>, <name> and <path> can be used.
type cmds [][]string

// Language represents a programming language with it Extension and Commands to
// execute its programs.
type Language struct {
	// Extension represents the file extension used by this language.
	Extension string
	// Commands  [][]string // placeholders: <name> file name (without
	// extension), <file> file name, <path> path without file name
	Commands cmds
}

// Supported Languages
const (
	Bash       = "bash"
	Zsh        = "zsh"
	Fish       = "fish"
	Elixir     = "elixir"
	Go         = "go"
	Javascript = "javascript"
	Lua        = "lua"
	OCaml      = "ocaml"
	Perl       = "perl"
	Python     = "python"
	Ruby       = "ruby"
	Rust       = "rust"
	Java       = "java"
	Julia      = "julia"
	C          = "c"
	Cpp        = "cpp"
	Swift      = "swift"
	Dart       = "dart"
	V          = "v"
	Scala      = "scala"
	Haskell    = "haskell"
)

// Languages is a map of supported languages with their extensions and commands
// to run to execute the program.
var Languages = map[string]Language{
	Bash: {
		Extension: "sh",
		Commands:  cmds{{"bash", "<file>"}},
	},
	Zsh: {
		Extension: "zsh",
		Commands:  cmds{{"zsh", "<file>"}},
	},
	Fish: {
		Extension: "fish",
		Commands:  cmds{{"fish", "<file>"}},
	},
	Elixir: {
		Extension: "exs",
		Commands:  cmds{{"elixir", "<file>"}},
	},
	Go: {
		Extension: "go",
		Commands:  cmds{{"go", "run", "<file>"}},
	},
	Javascript: {
		Extension: "js",
		Commands:  cmds{{"node", "<file>"}},
	},
	Lua: {
		Extension: "lua",
		Commands:  cmds{{"lua", "<file>"}},
	},
	Ruby: {
		Extension: "rb",
		Commands:  cmds{{"ruby", "<file>"}},
	},
	OCaml: {
		Extension: "ml",
		Commands:  cmds{{"ocaml", "<file>"}},
	},
	Python: {
		Extension: "py",
		Commands:  cmds{{"python", "<file>"}},
	},
	Perl: {
		Extension: "pl",
		Commands:  cmds{{"perl", "<file>"}},
	},
	Rust: {
		Extension: "rs",
		Commands: cmds{
			// compile code
			{"rustc", "<file>", "-o", "<path>/<name>.run"},
			// run compiled file
			{"<path>/<name>.run"},
		},
	},
	Java: {
		Extension: "java",
		Commands:  cmds{{"java", "<file>"}},
	},
	Julia: {
		Extension: "jl",
		Commands:  cmds{{"julia", "<file>"}},
	},
	Cpp: {
		Extension: "cpp",
		Commands: cmds{
			{"g++", "-std=c++20", "-o", "<path>/<name>.run", "<file>"},
			{"<path>/<name>.run"},
		},
	},
	C: {
		Extension: "c",
		Commands: cmds{
			{"gcc", "-o", "<path>/<name>.run", "<file>"},
			{"<path>/<name>.run"},
		},
	},
	Swift: {
		Extension: "swift",
		Commands:  cmds{{"swift", "<file>"}},
	},
	Dart: {
		Extension: "dart",
		Commands:  cmds{{"dart", "<file>"}},
	},
	V: {
		Extension: "v",
		Commands:  cmds{{"v", "run", "<file>"}},
	},
	Scala: {
		Extension: "sc",
		Commands: cmds{{"scala-cli", "run", "<file>"}},
	},
	Haskell: {
		Extension: "hs",
		Commands: cmds{{"runghc", "<file>"}},
	},
}
