package main
import ("errors";"fmt")
var ErrNotFound=errors.New("not found")
type VendorError struct{Code int}
func(e *VendorError)Error()string{return fmt.Sprintf("vendor code %d",e.Code)}
func vendorFetch(id int)error{if id==0{return &VendorError{404}};return &VendorError{503}}
func getUser(id int)error{err:=vendorFetch(id);var ve *VendorError;if errors.As(err,&ve)&&ve.Code==404{return fmt.Errorf("get user %d: %w",id,ErrNotFound)};return fmt.Errorf("get user %d: %w",id,err)}
func main(){fmt.Println(getUser(0))}

