#!/usr/bin/env bash
# @File:     commit.sh
# @Created:  2020-03-26 20:27:09
# @Modified: 2020-03-26 21:56:03
# @OA:       Antonio Escalera
# @CA:       Antonio Escalera
# @Mail:     aj@angelofdeauth.host
# @Copy:     Copyright © 2019 Antonio Escalera <aj@angelofdeauth.host>

set  -ex

cd "$(dirname "$0")/../"

git add .
git commit
git push

echo "[OK] Commit completed!"