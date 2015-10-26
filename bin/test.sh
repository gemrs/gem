#!/bin/bash

RETURN_CODE=0

test_package() {
    root_pkg=$(go list $1)
    root_profile=$2
    all_profiles=""
    for pkg in `go list ${root_pkg}/...`; do
        profile_path=$(echo $pkg | sed 's/\//_/g')
        profile_path+=".profile"
        go test -coverpkg=${root_pkg}/... -covermode=count -coverprofile=${profile_path} ${pkg}
        [ $? -ne 0 ] && RETURN_CODE=1
        [[ -f ${profile_path} ]] && all_profiles+=" ${profile_path}"
    done

    gocovmerge $all_profiles > $root_profile
    rm -f $all_profiles # Remove the old profiles to avoid confusion
}

test_package ./gem gem.profile
test_package ./bbc bbc.profile
gocovmerge gem.profile bbc.profile > coverage.profile
rm gem.profile bbc.profile # Remove the old profiles to avoid confusion
exit $RETURN_CODE
