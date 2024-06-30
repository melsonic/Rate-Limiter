package constants

import "fmt"

const PORT int = 3045
const BucketCapacity int = 10
const TokenRefillRate int = 1

// global values
var ServerAddr string = fmt.Sprintf(":%d", PORT)
