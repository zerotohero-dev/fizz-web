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

if [ -z "$IMAGE" ]
then
   echo "IMAGE is empty!"; exit 1;
fi

if [ -z "$TAG" ]
then
   echo "TAG is empty!"; exit 1;
fi

if [ -z "$REPO" ]
then
   echo "REPO is empty!"; exit 1;
fi

rm -rf ./usr/local/share

mkdir -p ./usr/local/share
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error creating usr/local/share."
  exit 1
fi

cp -R /usr/local/share/fizz ./usr/local/share
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Error copying the generated HTML."
  exit 1
fi

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
