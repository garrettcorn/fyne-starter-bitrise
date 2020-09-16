# fyne-starter-bitrise
fyne-starter-bitrise

# Setup
1. Clone this repo by using [this template](http://template.link "https://github.com/garrettcorn/fyne-starter-bitrise/generate")
1. Sign up at BitRise
    - [Referral Link](https://app.bitrise.io/referral/1cfba1cf5ffbaf60 "BitRise Referral Link") 
    - [Non-Referral Link](https://app.bitrise.io/ "BitRise Non-Referral Link")
1. [Add new App to BitRise](https://app.bitrise.io/apps/add "BitRise add app")
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
1. Click the "Register Webhook for me!" or "Skip the webhook registration" button.
1. A build will kick off, but it doesn't use the bitrise.yml file from the repo
1. Click "Open Workflow Editor"
1. Click "BitRise.yml"
1. Click "Store in app repository"
1. Click "Update Settings"
1. Click "Continue." Now the builds will be based off of the bitrise.yml file in the repository.
