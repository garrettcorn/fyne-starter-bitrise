# fyne-starter-bitrise
fyne-starter-bitrise is a quick start to CI for fyne. It provided the configuration necessary to start producting multi-platform and architecture applications using the go programming language.

# Setup
1. Clone this repo by using [the template link](https://github.com/garrettcorn/fyne-starter-bitrise/generate "template link")
1. Sign up at BitRise
    - [Referral Link](https://app.bitrise.io/referral/1cfba1cf5ffbaf60 "BitRise Referral Link") 
    - [Non-Referral Link](https://app.bitrise.io/ "BitRise Non-Referral Link")
1. [Add a new App to BitRise](https://app.bitrise.io/apps/add "BitRise add app")
    - Allow BitRise access to you github account if needed
1. Find this starter in your github repos
1. Type in "master" for branch selection
1. There will likely be a message about "An error occured during the validation process"
    - Select the "Restart scanning without validation" radio button
        - Click the "Next" button
1. Select "Other / Manual" for the project build configuration.
    - Chose a stack such as Android & Docker, on Ubuntu 16.04
        - Click the "I'm Ready" button
1. Upload an App Icon or click "Skip for Now"
1. Click the "Register Webhook for me!" button.
1. A build will kick off, but it doesn't use the bitrise.yml file from the repo
1. Click "Open Workflow Editor"
1. Click "BitRise.yml"
1. Click "Store in app repository"
1. Click "Update Settings"
1. Click "Continue." Now the builds will be based off of the bitrise.yml file in the repository.

# Result
When a change is pushed to the new repo a BitRise build will kick off and spit out Windows(amd64,386), Linux(amd64,386,arm,arm64), Darwin(amd64), and Android(combined apk) executables that can distributed and run on their respective platforms.

## Credit
This project was made possible by fyne and fyne-cross. Please take a look at their respective projects below.
- [fyne](https://fyne.io)
- [fyne-cross](https://github.com/lucor/fyne-cross)
