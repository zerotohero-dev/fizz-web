#!/usr/bin/env zsh

#  \
#  \\,
#   \\\,^,.,,.                     Zero to Hero
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

# shellcheck disable=SC1090
source ~/.zprofile

IMAGE=$ECR_IMAGE_FIZZ_WEB
TAG=$ECR_TAG_FIZZ_WEB
REPO=$ECR_REPO

# TODO: add more controls here.
rm -rf ./usr/local/share
mkdir -p ./usr/local/share
cp -R /usr/local/share/fizz ./usr/local/share

echo "»»» building"
docker build -t "$IMAGE":"$TAG" .
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error building the image."
  exit 1
fi

echo "»»» tagging"
docker tag "$IMAGE":"$TAG" "$REPO"/"$IMAGE":"$TAG"
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error tagging the image."
  exit 1
fi

echo "»»» pushing"
docker push "$REPO"/"$IMAGE":"$TAG"
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error pushing the image."
  exit 1
fi

echo "»»» Everything is awesome! «««"
