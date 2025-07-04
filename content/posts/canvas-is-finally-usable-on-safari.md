---
Slug: canvas-is-finally-usable-on-safari
Title: Canvas is finally usable on Safari
Summary: Safari has been a bit problematic when working with the <canvas> element, but after nearly a decade, Apple has finally improved the situation.
Date: 2025-05-10
Tags:
    - canvas
---

# Canvas is finally usable on Safari

Safari has been a bit problematic when it comes to the `<canvas>` element.
I first encountered this last year while troubleshooting a canvas editor web app.
The users, and more specifically iOS users, were unable to
export the results, which essentially rendered the app useless.
The issue ended up being that Safari simply could not handle a canvas larger
than **4096x4096** or a total area of **16,777,216** pixels.
If the limit is exceeded by even just a few pixels, the canvas will
break with the following warning in the console.

![Canvas area exceeds the maximum limit (width * height > 16777216).](/assets/images/posts/safari-canvas-error.png)

That's probably more than enough pixels for most use cases, but we found it really... *limiting*.
The editor canvas itself wasn't anywhere close to the limit,
but the exported result had to be scaled up and that's when we hit it.
While the original target resolution was on the larger side, at around **6000x6000**,
it's still a resolution you'd expect to be supported nowadays... or at least I did.

For comparison, Chrome on Android in 2018 set the limit to
**10,836x10,836** with a total area of **117,418,896** pixels.
I guess Apple just hasn't given much attention to this,
which is quite a shame considering that they **_invented_** the canvas element.
Unfortunately all you could do was go with a workaround like lowering the
resolution for iOS devices while waiting for Apple to improve the situation
somewhere in the future. In our case, **16,777,216** pixels was just barely enough
if we used every single one of them, so we went with the above-mentioned solution.

While I wasn't too happy about that, it was the best option at the moment.
I kept waiting and looking through Safari's release notes,
which often had some canvas related changes, but never the one I was looking for.

## Finally!

Fast forward to this year. I was working on a personal project involving canvas,
and while testing it on my iPhone, I noticed something had changed -
it was now outputting **5000x5000** sized images just fine.
Quick tests showed that the new limit was somewhere around **8000x8000**.
It was starting to look promising, but I still had to find
the exact limit as well as the iOS version where it was changed.
I checked the release notes again, but as expected, found no mention of this.

Next, I ran some tests using
<a href="https://www.npmjs.com/package/canvas-size" target="_blank" rel="noopener noreferrer">canvas-size</a>,
which showed that the new limit is **8192x8192** with a total area of **67,108,864** pixels.
Seeing that nothing had changed for my old iPad running iOS 17.7.5,
it seemed like they most likely changed it in iOS 18.0,
but I had to be confirm that another way, as my phone was running on a newer minor version.

So, I signed up for <a href="https://www.browserstack.com/"
target="_blank" rel="noopener noreferrer">BrowserStack</a> and ran the same tests on their
iPhone 13's with the following results:

| Version | Max Area    | Pixels       |
|---------|-------------|--------------|
| 17.7.2  | 4096 x 4096 | 16,777,216   |
| 18.0    | 8192 x 8192 | 67,108,864   |

Based on those test results, I think it's safe to say that
in **iOS 18.0** the `<canvas>` size limit has been increased to
much more generous **8192x8192** or a total area of **67,108,864** pixels.

I'm glad that this is finally solved. Unfortunately, the list of
issues and bugs in Safari is still looking just as long.
What makes these issues even more frustrating is the fact that
Apple has made it nearly impossible to debug Safari unless you own a Mac.
Even with an iPhone, I wasn't able to access the dev tools and had to
rely on BrowserStack to view the console.
BrowserStack saved the day here, but it's not really the best experience.
Next time I'll probably have to pay for it too, as I've pretty much burned
through the free credits doing this.

I just heard of another bug in Safari which has been quite problematic
for a business and creating a lot of work to get around it,
but I'll have to save that for another time.

## TLDR

For almost 10 years, the maximum canvas size on Safari was **4096x4096** (total: **16,777,216** pixels),
but in iOS 18.0, Apple quietly increased it to much more generous **8192x8192** (total: **67,108,864** pixels).
