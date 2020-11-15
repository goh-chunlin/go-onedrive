package main

import (
	"context"
	"fmt"

	"github.com/goh-chunlin/go-onedrive/onedrive"

	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: "EwBgA8l6BAAU6k7+XVQzkGyMv7VHB/h4cHbJYRAAAbe72vhP3r1FXWv1MGWFXB3gyDYCdnIyRvA6lmV4hcmlstjnxNINTVRFBdb9Lg0epzzfwxmsPg0xSycWkj0s6GTeZz+KmRQJsvpk3LkXccKWOK5MeKUn6v3Y9BSObhzW2VL6ZvUlNyDwNucMziAGNwCtmoadsCWrNAW+A1DZZAq9963Ii3/9EMoZYX6ZOgtOrGShiNbwuLt8LlEM2lYYHtyxQgyEQJ4kxanIKUQGd2v0iFgcX69Bt5ZXBMsUhDj2mLY2vgVhGyVh+dv0T0qx9I2ec2IddoX5hYiyLlKvjvAg7rslMF2gs6OOk62/cHKabARh0DOQBTbgjrf+Mr7H55wDZgAACGh+vEPygQ/PMALP/HUGHiByMo6wIykv7V6y90Pl0nK/9WSZppRu35DtIlCHuMu/Hv+YM7vdIHZQUEiogTw6asmE1kAK5t0s1jSPI9RS4lOqh/KseAp/nHG6qtV9ourr8fGP/jZrRcGXhhW5yc4zeSdFVYJ/sruEYURk+JXJyiZjd0cJ32Ui4RZhGPfEkjR5folYCnt+mxuz8DR3o+p9LVSGIW44506tdWcikXpsGaaWrjtHT4KJ3QOehpk+L4z9l1EuQyoISFKhA+7Tp725joC+t3aH8oiclJMk++kZDQRWoehzs4SZd2HJtxsStzamDGxPAI0TpZfQoh/u9v3gvzfSMZWDUym730FlrOWd/J51tEgNHl/Y9xB4iARRN7B3LcDTfe5Vq3/MYF7wpoCcfuuuY1nIraW1hIXzTMpc3S319htyVRen2svEdOOOTbDBPIVEGZXloWCB3vlc/BZOzDxJVMWh21Gc1nnbTwk1ImW1aFIhSpFwQGYJI7hprfztz5gE7jJvtWMmjwDPImkEdIX2Fms3NZMj1nM8ZszU0hoRIMfbOHfSiTeEq8vr20YJJvc6ewmtKvLxz6IG7DzYfXjn7uYrgfffqzMdh3J0OmRrXqnnMH6lPVW0QvMyqKzzHsFBPVDBOiZaW8DTVvN8UmmfZlt7yXEmm3nnQ3o04uo3ybnsXrCB47NA87DJfzSGhQGpQ2B27ef0fsfQOM+I80+z4WlHy+0x5/jGY/PUap6wccF7GAgywBbjRmQC"},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := onedrive.NewClient(tc)

	// list all drives for the authenticated user
	drives, err := client.Drives.List(ctx)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, drive := range drives.Drives {
		fmt.Printf("Results: %v\n", drive.Owner.User.DisplayName)
	}

	// list all drive items at the root for the authenticated user
	driveItems, err := client.DriveItems.List(ctx, "")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, driveItem := range driveItems.DriveItems {
		fmt.Printf("Results: %v\n", driveItem.Name)
	}

	// list all drive items at the Music folder for the authenticated user
	driveItems, err = client.DriveItems.List(ctx, "1F3A9FCD578789DD!2538")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	for _, driveItem := range driveItems.DriveItems {
		fmt.Printf("Results: %v\n", driveItem.Name)
	}
}
