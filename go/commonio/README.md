Reader Interface
Writer Interface

Types that implement Reader and Writer Interfaces
- os.File implements both io.Reader and io.Writer
- strings.Reader implements io.Reader
- bufio.Reader/Writer implements io.Reader and io.Writer respectively
- bytes.Buffer implements both io.Reader and io.Writer
bytes.Reader implements io.Reader
- Compress/gzip.Reader/Writer implements io.Reader and io.Writer respectively
- Crypto/cipher.StreamReader/StreamWriter implements io.Reader and io.Writer respectively
- Crypto/tls.Conn implements both io.Reader and io.Writer
- Encoding/csv.Reader/Writer implements io.Reader and io.Writer respectively
- Mime/multipart.Part implements io.Reader

ReaderAt and WriterAt Interfaces

ReadFrom and WriteTo Interfaces

Seeker interface

Closer interface

ByteReader and ByteWriter

ByteScanner, RuneReader, and RuneScanner

ReadCloser, ReadSeeker, ReadWriteCloser, ReadWriteSeeker, ReadWriter, WriteCloser, and WriteSeeker interfaces

SectionReader type
LimitedReader type
PipeReader and PipeWriter types

Copy and CopyN functions
ReadAtLeast and ReadFull functions
WriteString function
MultiReader and MultiWriter functions
TeeReader function
