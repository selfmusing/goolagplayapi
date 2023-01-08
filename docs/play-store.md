# Google Play Store

If you want API 24 or lower, you can just install user certificate and be done.
If you need API 25 or higher, first go here:

https://opengapps.org

I selected this option:

Platform | Android | Variant
---------|---------|--------
x86      | 6       | pico

but newer Android should work as well. Then open Android Studio. On "Select
Hardware" screen, select a device without "Play Store" icon.

## With Google APIs

On "System Image" screen, I selected this option:

API Level | ABI | Target
----------|-----|----------------------
24        | x86 | Android 7 Google APIs

but newer APIs should work as well. You should only need one file from the Zip
archive:

~~~
Core\vending-x86.tar.lz
~~~

Inside this will be another file:

~~~
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Now, start the device:

~~~
emulator -list-avds
emulator -avd Pixel_3a_XL_API_24 -writable-system
~~~

Next, install Google Play Store. Note that you cannot use the normal method of
drag APK to device screen, or you will get one of these errors:

~~~
The APK failed to install.<br/> Error: Could not parse error string

The APK failed to install.<br/> Error: INSTALL_FAILED_UPDATE_INCOMPATIBLE:
Package com.android.vending signatures do not match the previously installed
version; ignoring!

The APK failed to install.<br/> Error: INSTALL_PARSE_FAILED_NO_CERTIFICATES:
Failed to collect certificates from /data/app/vmdl1047870024.tmp/base.apk:
META-INF/BNDLTOOL.SF indicates /data/app/vmdl1047870024.tmp/base.apk is signed
using APK Signature Scheme v2, but no such signature was found. Signature
stripped?
~~~

Install like this:

~~~
adb root
adb remount
adb push Phonesky.apk /system/priv-app
adb reboot
~~~

After reboot, install system certificate, and start MITM program. You should
then be able to start Google Play Store as normal.

## Without Google APIs

Using the method above with Google APIs image, you still get some apps such as
YouTube. If you want to install different version of one of these apps, use
this method. On "System Image" screen, I selected this option:

API Level | ABI | Target
----------|-----|----------
24        | x86 | Android 7

but newer APIs should work as well. You need these files from the Zip archive:

~~~
Core\gmscore-x86.tar.lz
Core\vending-x86.tar.lz
~~~

Then extract these from the above files:

~~~
gmscore-x86\nodpi\priv-app\PrebuiltGmsCore\PrebuiltGmsCore.apk
vending-x86\nodpi\priv-app\Phonesky\Phonesky.apk
~~~

Use the same method above to install the APKs. After reboot, you should then be
able to install YouTube or whatever app. Note that unlike above, you dont
actually need to run the Google Play setup or even start the Google Play app at
the end.

- https://github.com/httptoolkit/httptoolkit/issues/200
- https://multigesture.net/articles/inspecting-https-network-traffic-of-any-android-app
