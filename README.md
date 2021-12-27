# wif - was it fast?

## Command line tool to measure and benchmark file transfer performance

While working on projects where I was uploading / downloading files I often found myself wondering about how long the upload / download alone _should_ be taking. I wanted to know this so I could reason about where potential bottlenecks in performance were coming from.

For example, when trying to understand performance of uploading a media file from a mobile device (raw file upload, file transcoding, converted file upload, converted file download) it would be useful to have an idea how long the upload and download should take depending on what type of network my mobile device is connected to.

### Example Usage

Question: Approximately how long should it take to upload `large_video_file.mp4` to an object store?

```
$ wif analyze large_video_file.mp4 -u
File Info
Name: large_video_file.mp4
Format: mp4
Size: 100mb

Benchmarks
Wi-Fi 2G (@ 20mbps): 20 seconds
Wi-Fi 5G (@ 50mbps): 5 seconds
5G wireless (@ 40mbps) : 7 seconds
4G LTE wireless (@ 30mbps): 10 seconds
3G wireless (@ 25mbps): 30 seconds
```

Question: Approximately how long should it take to download `large_video_file.mp4` from an object store?

```
$ wif analyze large_video_file.mp4 -d
File Info
Name: large_video_file.mp4
Format: mp4
Size: 100mb

Benchmarks
Wi-Fi 2G (@ 20mbps): 20 seconds
Wi-Fi 5G (@ 50mbps): 5 seconds
5G wireless (@ 40mbps) : 7 seconds
4G LTE wireless (@ 30mbps): 10 seconds
3G wireless (@ 25mbps): 30 seconds
```

Question: How long should it take to upload `large_video_file.mp4` to AWS us-east-1 from my current location?

```
$ wif analyze large_video_file.mp4 -u --target aws-east-1
File Info
Name: large_video_file.mp4
Format: mp4
Size: 100mb

Current location: New York, United States
Target location: Virginia, United States

Distance to target: 400 miles

Benchmarks
Wi-Fi 2G (@ 20mbps): 20 seconds
Wi-Fi 5G (@ 50mbps): 5 seconds
5G wireless (@ 40mbps) : 7 seconds
4G LTE wireless (@ 30mbps): 10 seconds
3G wireless (@ 25mbps): 30 seconds
```

Question: How long should it take to download `large_video_file.mp4` from AWS us-east-1 to my current location?

```
$ wif analyze large_video_file.mp4 -u --target aws-east-1
File Info
Name: large_video_file.mp4
Format: mp4
Size: 100mb

Current location:

Distance to us-e

Benchmarks
Wi-Fi 2G (@ 20mbps): 20 seconds
Wi-Fi 5G (@ 50mbps): 5 seconds
5G wireless (@ 40mbps) : 7 seconds
4G LTE wireless (@ 30mbps): 10 seconds
3G wireless (@ 25mbps): 30 seconds
```

Question: Approximately how long should it take to upload a 75mb file to an object store?

```
$ wif simulate 75mb -u
File Info
Name: large_video_file.mp4
Format: mp4
Size: 75mb

Benchmarks
Wi-Fi 2G (@ 20mbps): 20 seconds
Wi-Fi 5G (@ 50mbps): 5 seconds
5G wireless (@ 40mbps) : 7 seconds
4G LTE wireless (@ 30mbps): 10 seconds
3G wireless (@ 25mbps): 30 seconds
```

Question: Approximately how long should it take to download 75mb file from an object store?

```
$ wif simulate 75mb -d
File Info
Name: large_video_file.mp4
Format: mp4
Size: 100mb

Benchmarks
Wi-Fi 2G (@ 20mbps): 20 seconds
Wi-Fi 5G (@ 50mbps): 5 seconds
5G wireless (@ 40mbps) : 7 seconds
4G LTE wireless (@ 30mbps): 10 seconds
3G wireless (@ 25mbps): 30 seconds
```
