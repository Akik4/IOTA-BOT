clear
token=$1

if [ -z "$token" ] 
then
    echo "Please provide a token"
    read token
fi

echo "Token is $token"
go run . $token