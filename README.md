# I have moved to codeberg this project is here only because self update cannot handle the codeberg release functionality yet.

> "Don't be evil" — Google

Google's motto was originally "Don't be evil," but in 2018 they decided to remove it. Before doing so, they had expanded it to: "And remember... don't be evil, and if you see something that you think isn't right – speak up!" The irony of this is deeply felt, for even if one is not truly evil, one may nevertheless be perceived as such.

For years, Android developers have had to put up with Google's invasive telemetry and absurd user agreements that essentially require selling your soul just to use their Android developer IDE. Today that ends with **apknator**, I called it that because it sounds cool. The tool's main purpose is to compile and build `.apk` files and sign them. This is version 1, so I haven't added a distribution flag signer; it was too much work for a prototype.


# How to use

the basic loop is very simple, you need to run 6 commands to run your first application on an android phone assuming you have installed the discribution.

fist step is renaming the `ap_linux_amd64` to `ap`

next you run the following commands blindly and trust me.

first you make ap a cli command

`sudo mv ap /usr/local/bin/`

here we download the needed dependencies in ~/android-apk

`ap download`

we initializ the template and get into it

`ap template`

`cd my-app`

we compile the java code and it spits out an apk

`ap build`

for this step you need an android phone and you need it to be connected to your pc with developer mode turned on, and allow usb app downloadig

`adb install dist/unsigned_base-aligned-debugSigned.apk`