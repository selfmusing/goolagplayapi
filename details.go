package googleplay

import (
   "errors"
   "github.com/89z/rosso/protobuf"
   "github.com/89z/rosso/strconv"
   "io"
   "net/http"
   "net/url"
   "time"
)

func (d Details) MarshalText() ([]byte, error) {
   var b []byte
   b = append(b, "Title: "...)
   if val, err := d.Title(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nCreator: "...)
   if val, err := d.Creator(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nUpload Date: "...)
   if val, err := d.Upload_Date(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nVersion: "...)
   if val, err := d.Version(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   b = append(b, "\nVersion Code: "...)
   if val, err := d.Version_Code(); err != nil {
      return nil, err
   } else {
      b = strconv.AppendUint(b, val, 10)
   }
   b = append(b, "\nNum Downloads: "...)
   if val, err := d.Num_Downloads(); err != nil {
      return nil, err
   } else {
      b = strconv.AppendCardinal(b, val)
   }
   b = append(b, "\nInstallation Size: "...)
   if val, err := d.Installation_Size(); err != nil {
      return nil, err
   } else {
      b = strconv.AppendSize(b, val)
   }
   b = append(b, "\nFile:"...)
   for _, file := range d.File() {
      if val, err := file.File_Type(); err != nil {
         return nil, err
      } else if val >= 1 {
         b = append(b, " OBB"...)
      } else {
         b = append(b, " APK"...)
      }
   }
   b = append(b, "\nOffer: "...)
   if val, err := d.Micros(); err != nil {
      return nil, err
   } else {
      b = strconv.AppendUint(b, val, 10)
   }
   b = append(b, ' ')
   if val, err := d.Currency_Code(); err != nil {
      return nil, err
   } else {
      b = append(b, val...)
   }
   return append(b, '\n'), nil
}

func (d Details) Upload_Date() (string, error) {
   // .details.appDetails.uploadDate
   date, err := d.Get(13).Get(1).Get_String(16)
   if err != nil {
      return "", errors.New("uploadDate not found, try another platform")
   }
   return date, nil
}

func (h Header) Details(app string) (*Details, error) {
   req, err := http.NewRequest(
      "GET", "https://android.clients.google.com/fdfe/details", nil,
   )
   if err != nil {
      return nil, err
   }
   // half of the apps I test require User-Agent,
   // so just set it for all of them
   h.Set_Agent(req.Header)
   h.Set_Auth(req.Header)
   h.Set_Device(req.Header)
   req.URL.RawQuery = "doc=" + url.QueryEscape(app)
   res, err := Client.Do(req)
   if err != nil {
      return nil, err
   }
   defer res.Body.Close()
   body, err := io.ReadAll(res.Body)
   if err != nil {
      return nil, err
   }
   // ResponseWrapper
   response_wrapper, err := protobuf.Unmarshal(body)
   if err != nil {
      return nil, err
   }
   var det Details
   // .payload.detailsResponse.docV2
   det.Message = response_wrapper.Get(1).Get(2).Get(4)
   return &det, nil
}

type Details struct {
   protobuf.Message
}

// will fail with wrong ABI
func (d Details) Version_Code() (uint64, error) {
   // .details.appDetails.versionCode
   return d.Get(13).Get(1).Get_Varint(3)
}

// will fail with wrong ABI
func (d Details) Version() (string, error) {
   // .details.appDetails.versionString
   return d.Get(13).Get(1).Get_String(4)
}

// will fail with wrong ABI
func (d Details) Installation_Size() (uint64, error) {
   // .details.appDetails.installationSize
   return d.Get(13).Get(1).Get_Varint(9)
}

// should work with any ABI
func (d Details) Title() (string, error) {
   // .title
   return d.Get_String(5)
}

// should work with any ABI
func (d Details) Creator() (string, error) {
   // .creator
   return d.Get_String(6)
}

// should work with any ABI
func (d Details) Micros() (uint64, error) {
   // .offer.micros
   return d.Get(8).Get_Varint(1)
}

// should work with any ABI
func (d Details) Currency_Code() (string, error) {
   // .offer.currencyCode
   return d.Get(8).Get_String(2)
}

// should work with any ABI
func (d Details) Num_Downloads() (uint64, error) {
   // .details.appDetails
   // I dont know the name of field 70, but the similar field 13 is called
   // .numDownloads
   return d.Get(13).Get(1).Get_Varint(70)
}

// will fail with wrong ABI
func (d Details) File() []File_Metadata {
   var files []File_Metadata
   // .details.appDetails.file
   for _, file := range d.Get(13).Get(1).Get_Messages(17) {
      files = append(files, File_Metadata{file})
   }
   return files
}

// This only works with English. You can force English with:
// Accept-Language: en
func (d Details) Time() (time.Time, error) {
   date, err := d.Upload_Date()
   if err != nil {
      return time.Time{}, err
   }
   return time.Parse("Jan 2, 2006", date)
}
