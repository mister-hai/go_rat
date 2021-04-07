
WORK IN PROGRESS

This is a Remote Administration Tool
written in Golang

As a practice in learning Golang, reverse engineering, secure coding 
practices and post exploitation programming. The code will be drafted 
with a focus on clean, logically organized, well documentated structure.

    And as a tutorial for "the Church Of The Subhacker" Wiki, so it will be intentionally broken. You must learn, in order to use it.

This tutorial assumes some familiarity with programming concepts, languages, and networking

    This is not an introductory text

This project uses code from the following sources, please support every communities efforts to educate others by adding to the pool of available knowledge
     
     -https://github.com/bluesentinelsec/OffensiveGoLang.git

     -https://medium.com/@mat285/encrypting-streams-in-go-6cff6062a107


This code can be used to inadvertently break the law if you do not know what you are doing


Minor update:


    Use the python init/build script once it is finished, it will streamline the process



KNOWN GOOD DEBIAN VM CONFIGURATION:

        Download the Go language package for Linux
            - place in ~/ directory

        If you are creating a new project you make the following folder structure
            - everything code-y goes in /src
            - top level is for package/module/git stuff
            - name of package in all files needs to be the same

		Install the recomended go extension in VSCode
			- ctrl+shift+p - type in "Go : install/update tools"
			- select all the check boxes

		~/.bashrc:
            export GOROOT=/home/user/go
            export PATH=$GOROOT:$PATH:/home/user/.local/bin:/home/user/go/bin

		settings.json in VSCODE:

	{
    	"window.zoomLevel": 2,
    	"workbench.editorAssociations": [
        	{
            	"viewType": "jupyter.notebook.ipynb",
            	"filenamePattern": "*.ipynb"
        	}
    	],
    	"go.goroot": "/home/user/go",
    	"go.installDependenciesWhenBuilding": true,
    	"go.buildOnSave": "workspace",
    	"go.formatTool": "gofmt",
    	"go.languageServerFlags": [
        	"-rpc.trace"
      	]
	}

    For a new project:
	    Open the terminal in VSCode and type in, in the root directory of the package:
		    - go mod init go_rat

	THINGS SHOULD WORK DO NOT USE THE SUSPEND/SAVE STATE FUNCTION
	YOU WILL BREAK THE INSTALL (at least I did), just shut down the VM
	with the shutdown command

	... Careful changing the formatter to "gofmt" it hung my VM

    GITHUB USAGE:

        - github install on debian/linux
            sudo apt-get install git

        - downloading a project to hack on it
            git clone https://github.com/user_name/SomeOtherProject.git
        
        - git config 
            git config --global user.name "user_name"
            git config --global user.email "email_id"
            git remote add origin https://github.com/user_name/MyTestProject.git

        - git create new project
            git init MyTestProject

        - workflow to commit ALL files in project to github
         must be done after putting files in folder, be careful not to confuse github
         by swapping files around. if github isnt aware of specific changes it breaks stuff
         just avoid using the web interface to modify files in the repository on important projects until you understand how github expects changes to occur.
         Its ok to download packages from the web interface but you should clone them
         with "git clone" in the terminal
            git add .
            git commit -m "some_message_NARF!"
            git push origin master (OR main OR whatever the branch name is)

	USEFUL GO COMMANDS:
        go get github.com/user/package
            will download and add new stuff


	USEFUL GO TIPS:
		-DO NOT USE existing names in modules PERIOD, you will confuse the linter/compiler
			e.g. "color" is a module name so you cant use "color" as a variable name

Every project I start turns into something enourmous... why do I do these things to myself?