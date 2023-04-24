# If you come from bash you might have to change your $PATH.
# export PATH=$HOME/bin:/usr/local/bin:$PATH

# Path to your oh-my-zsh installation.
export ZSH=$HOME/.oh-my-zsh

# Set name of the theme to load --- if set to "random", it will
# load a random theme each time oh-my-zsh is loaded, in which case,
# to know which specific one was loaded, run: echo $RANDOM_THEME
# See https://github.com/ohmyzsh/ohmyzsh/wiki/Themes
ZSH_THEME="af-magic"

# Set list of themes to pick from when loading at random
# Setting this variable when ZSH_THEME=random will cause zsh to load
# a theme from this variable instead of looking in $ZSH/themes/
# If set to an empty array, this variable will have no effect.
# ZSH_THEME_RANDOM_CANDIDATES=( "robbyrussell" "agnoster" )

# Uncomment the following line to use case-sensitive completion.
# CASE_SENSITIVE="true"

# Uncomment the following line to use hyphen-insensitive completion.
# Case-sensitive completion must be off. _ and - will be interchangeable.
# HYPHEN_INSENSITIVE="true"

# Uncomment one of the following lines to change the auto-update behavior
# zstyle ':omz:update' mode disabled  # disable automatic updates
# zstyle ':omz:update' mode auto      # update automatically without asking
# zstyle ':omz:update' mode reminder  # just remind me to update when it's time

# Uncomment the following line to change how often to auto-update (in days).
# zstyle ':omz:update' frequency 13

# Uncomment the following line if pasting URLs and other text is messed up.
# DISABLE_MAGIC_FUNCTIONS="true"

# Uncomment the following line to disable colors in ls.
# DISABLE_LS_COLORS="true"

# Uncomment the following line to disable auto-setting terminal title.
# DISABLE_AUTO_TITLE="true"

# Uncomment the following line to enable command auto-correction.
# ENABLE_CORRECTION="true"

# Uncomment the following line to display red dots whilst waiting for completion.
# You can also set it to another string to have that shown instead of the default red dots.
# e.g. COMPLETION_WAITING_DOTS="%F{yellow}waiting...%f"
# Caution: this setting can cause issues with multiline prompts in zsh < 5.7.1 (see #5765)
# COMPLETION_WAITING_DOTS="true"

# Uncomment the following line if you want to disable marking untracked files
# under VCS as dirty. This makes repository status check for large repositories
# much, much faster.
# DISABLE_UNTRACKED_FILES_DIRTY="true"

# Uncomment the following line if you want to change the command execution time
# stamp shown in the history command output.
# You can set one of the optional three formats:
# "mm/dd/yyyy"|"dd.mm.yyyy"|"yyyy-mm-dd"
# or set a custom format using the strftime function format specifications,
# see 'man strftime' for details.
# HIST_STAMPS="mm/dd/yyyy"

# Would you like to use another custom folder than $ZSH/custom?
# ZSH_CUSTOM=/path/to/new-custom-folder

# Which plugins would you like to load?
# Standard plugins can be found in $ZSH/plugins/
# Custom plugins may be added to $ZSH_CUSTOM/plugins/
# Example format: plugins=(rails git textmate ruby lighthouse)
# Add wisely, as too many plugins slow down shell startup.
plugins=(git)

source $ZSH/oh-my-zsh.sh

# User configuration

# export MANPATH="/usr/local/man:$MANPATH"

# You may need to manually set your language environment
# export LANG=en_US.UTF-8

# Preferred editor for local and remote sessions
# if [[ -n $SSH_CONNECTION ]]; then
#   export EDITOR='vim'
# else
#   export EDITOR='mvim'
# fi

# Compilation flags
# export ARCHFLAGS="-arch x86_64"

# Set personal aliases, overriding those provided by oh-my-zsh libs,
# plugins, and themes. Aliases can be placed here, though oh-my-zsh
# users are encouraged to define aliases within the ZSH_CUSTOM folder.
# For a full list of active aliases, run `alias`.
#
# Example aliases
# alias zshconfig="mate ~/.zshrc"
# alias ohmyzsh="mate ~/.oh-my-zsh"


export GOSUMDB="off"
# export GOPROXY="http://goproxy.intra.xiaojukeji.com"
#export PATH=/usr/local/go1.14.1/bin:$PATH
#export PATH=/usr/local/go1.18.9/bin:$PATH
export PATH=/usr/local/go1.20/bin:$PATH
export PATH=/Users/ethan/go/bin:$PATH

#设置时区数据库
export ICU_TIMEZONE_FILES_DIR=/Users/didi/elvish/icu_data/latest/le
#设置ICU的LIB库地址
export DYLD_LIBRARY_PATH=/Users/didi/elvish/icu4c/icu_Darwin/lib
export CGO_CPPFLAGS="-I/Users/didi/elvish/icu4c/icu_Darwin/include"
export CGO_LDFLAGS="-L/Users/didi/elvish/icu4c/icu_Darwin/lib"

export TB="ssh ethanwangyujie@relay.sys.xiaojukeji.com -p 35000"
alias tb="sshpass -p 'Wyj19930316.,' ssh ethanwangyujie@relay.sys.xiaojukeji.com -p 35000"
alias tp="sshpass -p 'Wyj19930316.,' sftp -oPort=8000 ethanwangyujie@fswap.sys.xiaojukeji.com"
alias ethan_dev="sshpass -p 'n~Y,nORz8r9s@ac' ssh root@10.96.83.172"
# export PATH=/Users/didi/redis-6.2.6/src:$PATH
# export PATH=/Users/local/bin:$PATH
# export GOPATH=/Users/didi/go
LANG=zh_CN.UTF-8
export LANG
export LC_ALL=en_US.UTF-8

[[ -s `brew --prefix`/etc/autojump.sh ]] && . `brew --prefix`/etc/autojump.sh

### micous config ###
export GOPRIVATE=gitlab-vywrajy.micoworld.net 
export GOPROXY=https://proxy.golang.org,direct
proto_generate='protoc -I ./proto -I ./proto_inner/ --go_out=plugins=grpc,,paths=source_relative:proto/pb/ --grpc-gateway_out=paths=source_relative:proto/pb/  proto_inner/*.proto proto/*.proto'

# 跳板机
alias mtb="ssh -p 2222 yujie@jumps.micoworld.net"

