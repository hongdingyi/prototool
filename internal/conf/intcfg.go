package conf

var (
	tmpl = `###################基础配置#####################
#导入路径可以设置环境变量: $GOPATH 或者 ${GOPATH}
#项目基础导入目录
import_path: $IDL_PATH

#当前项目依赖的proto文件
protos:
  - usermgt/passport/passport.proto
  - usermgt/code.proto
  - usermgt/user.proto
#依赖导入目录
includes:
  - $GOPATH/src/github.com/bilibili/kratos/third_party


####################lint配置####################
lint:
  # The lint group to use.
  # Available groups: "uber1", "uber2", "google", "empty".
  # The default group is the "uber1" lint group for backwards compatibility reasons,
  # however we recommend using the "uber2" lint group.
  # The special group "empty" has no linters, allowing you to manually specify all
  # lint rules in lint.rules.add.
  # Run prototool lint --list-all-lint-groups to see all available lint groups.
  # Run prototool lint --list-lint-group GROUP to list the linters in the given lint group.
  #group: google

  # Linter files to ignore.
  # These can either be file or directory names.
  # If there is a directory name, that directory and all sub-directories will be ignored.
  #ignores:
  #  - id: RPC_NAMES_CAMEL_CASE
  #      files:
  #        - path/to/foo.proto
  #        - path/to/bar.proto
  #  - id: SYNTAX_PROTO3
  #    files:
  #        - path/to/dir

  # Linter rules.
  # Run prototool lint --list-all-linters to see all available linters.
  # Run prototool lint --list-linters to see the currently configured linters.
  rules:
    add:
      - ENUM_NAMES_CAMEL_CASE
      - ENUM_NAMES_CAPITALIZED
    remove:
      #- ENUM_NAMES_CAMEL_CASE

####################编译配置####################
generate:
    go_options:
      extra_modifiers:
    plugins:
      - name: gofast
        flags: plugins=grpc
        output: ./genproto
`
)
