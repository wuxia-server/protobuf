for dirname in "message" ;
do
    if [ $dirname = "google" ]; then
        echo "skip google"
    else
        echo "deal with $dirname"
        filename="./golang/$dirname/$dirname.uri.go"
        `echo "package $dirname" > $filename`
        `echo "" >> $filename`
        for fileName in `ls ./pb/$dirname/*.proto` ;
        do
          msgName=`cat $fileName | egrep -o "^message (C2S|S2C|S2S).*{" | awk -F " |{" '{print $2}'`
          fileNoSuffix=$(basename ${fileName} .proto)
          for msg in $msgName
          do
              uri=`echo $msg| tr 'A-Z' 'a-z'`
              uriName=${msg//_/}
              `echo "const $uriName = \"/$fileNoSuffix/$uri/\"" >> $filename`
          done
        done
        echo "type uri string"

    fi
done


