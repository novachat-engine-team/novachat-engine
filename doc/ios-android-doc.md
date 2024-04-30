

https://core.telegram.org/reproducible-builds

Reproducible Builds for iOS and Android
This page contains instructions for verifying that Telegram's open source code is exactly the same as the code that is used to build the apps that are available in the App Store, Google Play and directly on the Telegram website.

Warning: Telegram supports reproducible builds as of version 5.13. Bear in mind that, at this stage, the verification process should be considered experimental. We will be updating our apps and these instructions to make this process as straightforward as possible.

Telegram for Android
Telegram for iOS
Please read the relevant notes and troubleshooting section carefully.

Dude in a jacket inspecting the hologram of a mechanical dog to verify it's built according to the blueprints provided.
Reproducible Builds for Android
Step 1. Install Docker
Docker can be obtained here. Once the installation is complete, log into your Docker account > Settings > Resources > Advanced and configure the amount of resources Docker may use:

Docker Performance
Docker Performance

We recommend using the maximum amount allowed by your system's hardware, in order to speed up the build time.

Step 2. Confirm which version you have installed on your Android device
You can find the version/build number and the source (website, Play Store, Huawei Store) at the bottom of the Settings page. Note that Telegram supports reproducible builds starting with version 5.13.

App Version
App Version

The commit tag to checkout source code for the example above will be release-9.3.3_3026.

Please make sure that you're using the correct version and build number of the version you want to check (and not the one from this example 😈).

Step 3. Obtain the source code
Open Terminal, run the commands:
git clone https://github.com/DrKLO/Telegram.git $HOME/telegram-android
cd $HOME/telegram-android
git checkout release-{VERSION AND BUILD NUMBER FROM STEP 2}

For our example, the command would be:
git checkout release-5.13.0_1821

Obtaining the source code
Step 4. Build the app
Open Terminal, run the commands:
cd $HOME/telegram-android
docker build -t telegram-build .

Building the app
docker run --rm -v "$PWD":/home/source telegram-build

Building the app
Building the app
These commands will produce 3 different APKs and 2 bundles:

/apk/afat/standalone/app.apk – used for direct downloads from telegram.org/android
/apk/afat/elease/app.apk – the playstore version
/apk/afat/release/app-huawei.apk – used exclusively for the Huawei store

bundle/bundleAfat_SDK23Release/TMessagesProj_App-bundleAfat_SDK23-release.aab
bundle/bundleAfatRelease/TMessagesProj_App-bundleAfat-release.aab

These APKs can be found in:
$HOME/telegram-android/TMessagesProj/build/outputs/apk/afat/

Use the folder name from Step 4 to find the correct folder that holds the same APK as installed on your device. For example, for the Play Store version, the path to your APK will be:

$HOME/telegram-android/TMessagesProj/build/outputs/apk/afat/release/app.apk
Copy this APK to the root source directory by running this command in Terminal:
cp $HOME/telegram-android/TMessagesProj/build/outputs/apk/afat/release/app.apk $HOME/telegram-android/telegram_built.apk

Copy the APK
Step 5. The Telegram APK installed on your device
You will need adb for this step.

ADB
ADB

If you downloaded your APK directly from Telegram's website, use the package name org.telegram.messenger.web in this step. To verify the Google Play APK, use org.telegram.messenger.

Connect your device to the computer, open Terminal, run the commands:
adb shell pm path org.telegram.messenger

ADB
The output will look something like this:
package:/data/app/org.telegram.messenger-_zOSURFEx2GpHM8UDF_PVg==/base.apk
By using this information, pull the APK from your device to $HOME/telegram-android using command:
adb pull /data/app/org.telegram.messenger-_zOSURFEx2GpHM8UDF_PVg==/base.apk $HOME/telegram-android/telegram_store.apk

Step 6. Compare the APKs
To compare Direct and Huawei Store versions, open Terminal, run the commands:
cd $HOME/telegram-android
python apkdiff.py telegram_store.apk telegram_built.apk
If the APKs are the same, you will see
APKs are the same!

Match
Play Store versions built from a bundle will be marked as 'store bundled'. To verify such builds, use:

python apkfrombundle.py telegram_bundle.aab telegram_store.apk

If anything goes wrong, you will see this:

Mismatch
If your APKs don't match, please make sure that you chose the correct code version and the right SDK.

Check out the Troubleshooting section first in case you run into trouble.

Reproducible Builds for iOS
The verification process for iOS builds is, unfortunately, a lot more complex than for Android. The two main issues with Apple's current policies and infrastructure are as follows:

Apple insists on using FairPlay encryption to “protect” even free apps from “app pirates” which makes obtaining the executable code of apps impossible without a jailbroken device. To solve this issue, Apple would simply need to allow submitting unencryptable binaries to the App Store. This would not affect security since the code would still be signed – and would enable anyone to check the integrity of apps supporting reproducible builds without endangering the integrity and security of their devices.

Building your own reproducible binaries is difficult because macOS doesn't support containers like Docker. If Apple followed in the footsteps of Linux (and even Microsoft!) and added container support, it would eliminate the need for steps 1-3 in the guide below.

As things stand now, you'll need a jailbroken device, at least 1,5 hours and approximately 90GB of free space to properly set up a virtual machine for the verification process.

To provide a stable and easily reproducible environment, Telegram iOS builds are compiled on a virtual machine. Parallels is used to verify the builds.

Despite the compiler bugs introduced by Apple in Xcode 14 (read more), we were able to restore deterministic builds using manually crafted linker flags. Use the steps below to verify builds compiled with XCode 13 and below, see here for XCode 14 instructions.

Step 1. Install the Parallels virtual machine
Parallels can be obtained here, it features a fully-functional trial version.

Step 2. Install the latest version of macOS Big Sur
To download an image that can be installed on the virtual machine, open the App Store, search for “Catalina” and click “View”.

Search for macOS Catalina on App Store > View
Search for macOS Catalina on App Store > View

macOS Catalina > Get
macOS Catalina > Get

This will open a system pop-up offering to download the OS:

macOS Catalina > Download
Choose “Download” and wait for the download to finish. > If you were not using the latest version of the OS, your system may start updating instead. Please finish updating to download macOS Catalina. When done, open Parallels and choose macOS Catalina:
Select 'Install Windows or another OS' > Continue
Select 'Install Windows or another OS' > Continue

Select a file... > Applications (All files) > Install macOS Сatalina
Select a file... > Applications (All files) > Install macOS Сatalina

Before starting the installation, configure the virtual machine:

Enable checkbox 'Customize settings before installation'
Checkbox 'Customize settings before installation'

Change the name of the virtual machine to macos11_Xcode12_5_1

Name VM as macos10_15_Xcode11_2
Name VM as macos10_15_Xcode12_2

Hardware > Processors: 2-4
Memory > 4GB may suffice but 8GB is recommended

At least 2 CPUs + 4 (8 recommended) GB Memory.
At least 2 CPUs + 4 (8 recommended) GB Memory.

You will get something like this:

Click Continue
Click Continue

Parallels may request access to your microphone and camera, this is not required – just press Close.

Install macOS > Continue
Install macOS > Continue

Your Apple ID is also not required, you can choose Set Up Later.

Skip Apple ID with 'Set Up Later'
Skip Apple ID with 'Set Up Later'

Use “telegram” for both the account name and password.

Do not ever use the password “telegram” for anything else, it's cursed.

Create a computer account with 'telegram' set both as account name and password
Create a computer account with 'telegram' set both as account name and password

Now install Parallels tools from the menu bar:

Install Parallels Tools using menu bar > Parallels icon > Actions > Install (Reinstall) Parallels Tools Install Parallels Tools
Install Parallels Tools using menu bar > Parallels icon > Actions > Install (Reinstall) Parallels Tools...

After the system restarts, log in.
Open Terminal and run:
sudo visudo
Enter the password “telegram”

Run Terminal on the VM, enter 'sudo visudo' > enter password (telegram)
Find this line at the end of the file:
%admin ALL=(ALL) ALL
Press “i” on your keyboard, add “NOPASSWD:”
%admin ALL=(ALL) NOPASSWD: ALL
Press Escape.
Type in “:wq”
Press Enter

Press i to edit the highlighted string
Press i to edit the highlighted string.

Enter :wq > press Enter
Enter :wq > press Enter.

In the terminal, run:
sudo systemsetup -setcomputersleep Never

sudo systemsetup -setcomputersleep Never > press Enter
sudo systemsetup -setcomputersleep Never > press Enter.

Step 3. Install SSH keys on the virtual machine
In the virtual machine, open System Settings > Sharing and enable Remote Login.

Go to macOS Settings > Sharing > enable Remote Login
In the virtual machine, open Terminal and run:
mkdir -p .ssh; nano .ssh/authorized_keys

Enter mkdir -p .ssh; nano .ssh/authorized_keys > press Enter
In your main OS, open Terminal and run:
if [ ! -e ~/.ssh/id_rsa.pub ]; then ssh-keygen -t rsa -b 4096; fi && cat ~/.ssh/id_rsa.pub | pbcopy

Terminal
If you see the line “Enter file in which to save the key (/Users/…/.ssh/id_rsa):”, press Enter
In the virtual machine, press CMD+V
Then Ctrl+O, Ctrl+X

Press Cmd+V > Ctrl+O > Ctrl+X.
Step 4. Install Xcode version 12.5.1
In the virtual machine, open Safari and go to https://developer.apple.com
Sign in to your Account:

developer.apple.com > Account > sign in with your Apple ID
developer.apple.com > Account > sign in with your Apple ID

Go to Downloads > More
Enter Xcode in the search field and find the version 12.5.1

Downloads > More > 12.5.1
Downloads > More > Xcode 12.5.1

Once the installation is complete, open the file Xcode 12.5.1.xip. The system will unarchive the app into the same folder. Move it to the Applications folder using Finder.

Unarchive Xcode > drag the app to Applications folder
Unarchive Xcode > drag the app to Applications folder

Move Xcode to Applications using Finder > Run > Agree to install additional components
On the virtual machine, run this command from the terminal:
sudo xcode-select -s /Applications/Xcode.app/Contents/Developer

sudo xcode-select -s /Applications/Xcode.app/Contents/Developer
Shut down the virtual machine.

Shut down the virtual machine
Shut down the virtual machine

Step 4.1
Download the certificates at https://github.com/TelegramMessenger/Telegram-iOS/tree/master/build-system/fake-codesigning/certs/distribution and install them into the virtual machine.

Launch Keychain Access and double-click the installed certificate. Under “Trust”, change “When using this certificate” to “Always Trust”.

Step 5. Obtaining the source code
git clone --recursive https://github.com/TelegramMessenger/telegram-ios.git $HOME/telegram-ios
cd $HOME/telegram-ios
git checkout release-${VERSION_NUMBER}

E.g., git checkout release-7.3. Please note that you need to check out the whole git history as the build version depends on the number of commits in the repository.

git clone
Step 6. Downloading Bazel 3.7.0 to $HOME/bazel/bazel
mkdir -p $HOME/bazel && cd $HOME/bazel
curl -O -L https://github.com/bazelbuild/bazel/releases/download/3.7.0/bazel-3.7.0-darwin-x86_64
mv bazel-3.7.0-darwin-x86_64 bazel

Check that you have downloaded the correct version:
chmod +x bazel
./bazel --version

Step 7. Building the app
Open Terminal, run the commands:
cd $HOME/telegram-ios BAZEL="$HOME/bazel/bazel" sh buildbox/build-telegram.sh verify

Building the app
If the environment has been set up correctly, this will start the building process. Note that this step can easily take 30-40 minutes. The average build time on a MacBook Pro (i9 6 core) is 35 minutes.

Building started Building completed
Once the process is complete the resulting IPA file can be found in build/artifacts/Telegram.ipa
All the following steps will be made via Terminal on your main system.

Step 8. Downloading a decrypted version of the app from the App Store
This step requires a jailbroken device equipped with tools for decrypting apps. We‘d love to make this process more simple but that’s what you get for using Apple tech.

Step 9. Comparing the AppStore build and the version built in the virtual machine
Install the necessary tools:
if ! type brew > /dev/null; then /usr/bin/ruby -e "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install)"; fi && brew install python3

Installing python
Run
python3 tools/ipadiff.py build/artifacts/Telegram.ipa PATH-TO-THE-IPA-FILE-FROM-STEP-9

cd telegram-ios > Enter
python3 tools/ipadiff.py build/artifacts/Telegram.ipa /path/ > Enter
Сomparing builds
In case of a successful comparison, you will get a text along these lines:

IPAs are equal, except for the files that can't currently be checked:
Excluded files that couldn't be checked due to being encrypted:
PlugIns/SiriIntents.appex/SiriIntents
PlugIns/Widget.appex/Widget
PlugIns/NotificationContent.appex/NotificationContent
PlugIns/NotificationService.appex/NotificationService
PlugIns/Share.appex/Share
IPAs contain Watch directory with a Watch app which can't be checked currently.
IPAs contain .car (Asset Catalog) files that are compiled by the App Store and can't currently be checked:

        Frameworks/TelegramUI.framework/Assets.car
        Assets.car
    IPAs contain .nib (compiled Interface Builder) files that are compiled by the App Store and can't currently be checked:
        Base.lproj/LaunchScreen.nib
The result > equal IPAs
In case of any mismatches, you'll get a detailed report.

Mismatching IPAs
iOS, XCode 14
Due to compiler bugs introduced by Apple in Xcode 14 (read more), you will need to use the modified instructions below to verify the latest builds:

Remove steps 6, 7

Steps 1-4.1 are replaced with:

Running Darwin-Containers:
Checkout the latest DarwinContainers code: git clone https://github.com/ali-fareed/darwin-containers.git
Open darwin-containers/DarwinContainers.xcodeproj
Select DarwinContainersDaemon target
In Signing & Capabilities select your team and set a unique bundle id
Run
Creating an OS image:
./darwin-containers fetch

Download the appropriate macOS restore image (e.g. 13.0):

./darwin-containers fetch "macOS 13.0"

Create a new OS image:

./darwin-containers create --source "macOS 13.0" --tag "macos-13.0-xcode-14.1" --manual

Follow the installation instructions. Set username to containerhost and password to containerhost.
Enable Remote Login and allow full disk access for remote users.
Connect to the guest VM using SSH with username containerhost and password containerhost.
Create directory ~/.ssh and set up the authorized_keys using the public key string printed by the darwin-containers create command earlier.
Upload the appropriate version of Xcode via scp and install to /Applications. Run it at least once to complete installation.
Shutdown the guest OS.

Obtaining verification IPA:
python3 -u build-system/Make/Make.py remote-build --darwinContainers="path-to-darwin-containers-script" --darwinContainersHost="unix://$HOME/.darwin-containers.sock" --configurationPath="build-system/appstore-configuration.json" --codesigningInformationPath=build-system/fake-codesigning --configuration=release_arm64

For more information see:

build-system/Make/RemoteBuild.py
.gitlab-ci.yml lane verify_beta_testflight

iOS: Notes
You will get a warning if the archive created in Step 7 contains encrypted files. If all these files are in the PlugIns subfolder, they represent various system extensions (e.g. external sharing, Siri, 3D touch). Decrypting such files using existing ways of receiving app archives via Jailbreak is non-trivial (but we're working on resolving this issue). If you do manage to decrypt them, e.g. on iOS 8, they will be matched.

You will be notified if the archive includes an Apple Watch app. The watch app will soon no longer be included in the archive.

Files with the .car extension are app resource archives (images, sounds) which were compiled and processed specifically for the target device. The App Store processes them in non-trivial ways, we're planning on getting rid of them in future versions.

The LaunchScreen.nib file is an empty file containing a description of the interface which is displayed by the system before the app is launched. It is processed by the App Store in a non-trivial way but doesn't contain any code and therefore may be ignored.

Troubleshooting
If you encounter any issues with obtaining the code, building and comparing the apps, please contact us at @BotSupport and include the hashtag #reproducibleBuilds with your message describing the problem.

Troubleshooting: Android
Make sure that you checkout the correct version of the code.

Make sure that you build the app using the right SDK.

If the gradle version used in the Dockerfile is not available anymore and building of the Docker image fails, wait for a Dockerfile update or update manually to lastest available version.

Troubleshooting: iOS
UPD: Despite the fact that the issue below persists, we were able to restore deterministic builds using manually crafted linker flags. See these updated steps for XCode 14.

Due to recent changes introduced in XCode 14 by Apple, it is currently not possible to create reproducible builds for iOS using tools officially supported by Apple. We will update this page as soon as Apple resolves the issue.

To confirm the issue for yourself, follow these steps:

Unzip link-issue.zip
sh test-link.sh
Compare link1.output and link.output. Specifically:
With some probability, the ordering of the LC_LOAD_DYLIB commands will vary.
The __LIKEDIT section will always vary.