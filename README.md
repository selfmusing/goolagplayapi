# goolagplayapi

> I’m not really sure what the point of this video is, but I guess just be
> generous.
>
> Be kind to people, because you never know how much they might need it, or how
> far it’ll go.
>
> [NakeyJakey (2018)](//youtube.com/watch?v=Cr0UYNKmrUs)

Download APK from Google Play or send API requests

## How to install?

This module works with Windows, macOS or Linux. First, [download Go][2] and
extract archive. Then [download&nbsp;GooglePlay][3] and extract archive. Then
navigate to `googleplay-master/cmd/googleplay`, and enter:

~~~
go build
~~~

[2]://go.dev/dl
[3]://github.com/89z/googleplay/archive/refs/heads/master.zip

## Tool examples

Before trying these examples, make sure the Google account you are using has
logged into the Play&nbsp;Store at least once before. Also you need to have
accepted the Google Play terms and conditions. Create a file containing token
(`aas_et`) for future requests:

~~~
googleplay -email EMAIL -password PASSWORD
~~~

Create a file containing `X-DFE-Device-ID` (GSF ID) for future requests:

~~~
googleplay -device
~~~

Get app details:

~~~
> googleplay -a com.google.android.youtube
Title: YouTube
Creator: Google LLC
Upload Date: Aug 24, 2022
Version: 17.33.42
Version Code: 1531188672
Num Downloads: 12.45 billion
Installation Size: 48.99 megabyte
File: APK APK APK APK
Offer: 0 USD
~~~

Purchase app. Only needs to be done once per Google account:

~~~
googleplay -a com.google.android.youtube -purchase
~~~

Download APK. You need to specify any valid version code. The latest code is
provided by the previous details command. If APK is split, all pieces will be
downloaded:

~~~
googleplay -a com.google.android.youtube -v 1531188672
~~~


Software is not licensed for commercial use. If you wish to purchase a
commercial license, or for other business questions, contact me:

- srpen6@gmail.com
- Discord srpen6#6983
