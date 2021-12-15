#!/usr/bin/env bash

package=$1
if [[ -z "$package" ]]; then
  echo "usage: $0 <package-name>"
  exit 1
fi
package_split=(${package//\// })
package_name=${package_split[-1]}

platforms=("windows/amd64" "windows/386" "linux/amd64" "linux/386")

for platform in "${platforms[@]}"
do
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    output_name=$package_name'-'$GOOS'-'$GOARCH
    if [ $GOOS = "windows" ]; then
        output_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $output_name $package
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
done
mv SnowPearDNS-linux-386 ./release/snowpear_32
mv SnowPearDNS-linux-amd64 ./release/snowpear_64
mv SnowPearDNS-windows-386.exe ./release/snowpeardns_32.exe
mv SnowPearDNS-windows-amd64.exe ./release/snowpeardns_64.exe