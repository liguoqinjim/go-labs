# 得到图片，不是通过取得src的方式
下载网页里面的图片其实很简单，我们可以得到img标签的src属性，然后直接下载就可以了。
但是有的时候，我们需要截图验证码，验证码图片的src，每次访问返回的可能都是一张新的图片，所以我们就不能使用这种方式。
那我们就采取，直接截屏，然后再截屏的图片里面裁切出验证码的图片。
