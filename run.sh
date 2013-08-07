#!/bin/sh
SCRIPTPATH=`dirname "$0"`
chmod u+x "$SCRIPTPATH/myapp"
"$SCRIPTPATH/myapp" -importPath myapp -srcPath "$SCRIPTPATH/src" -runMode prod
