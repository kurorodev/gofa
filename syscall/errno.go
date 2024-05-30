package syscall

const (
	/* From <errno.h>. */
	EPERM   = 1  /* Operation not permitted */
	ENOENT  = 2  /* No such file or directory */
	ESRCH   = 3  /* No such process */
	EINTR   = 4  /* Interrupted system call */
	EIO     = 5  /* Input/output error */
	ENXIO   = 6  /* Device not configured */
	E2BIG   = 7  /* Argument list too long */
	ENOEXEC = 8  /* Exec format error */
	EBADF   = 9  /* Bad file descriptor */
	ECHILD  = 10 /* No child processes */
	EDEADLK = 11 /* Resource deadlock avoided */
	/* 11 was EAGAIN */
	ENOMEM  = 12 /* Cannot allocate memory */
	EACCES  = 13 /* Permission denied */
	EFAULT  = 14 /* Bad address */
	ENOTBLK = 15 /* Block device required */
	EBUSY   = 16 /* Device busy */
	EEXIST  = 17 /* File exists */
	EXDEV   = 18 /* Cross-device link */
	ENODEV  = 19 /* Operation not supported by device */
	ENOTDIR = 20 /* Not a directory */
	EISDIR  = 21 /* Is a directory */
	EINVAL  = 22 /* Invalid argument */
	ENFILE  = 23 /* Too many open files in system */
	EMFILE  = 24 /* Too many open files */
	ENOTTY  = 25 /* Inappropriate ioctl for device */
	ETXTBSY = 26 /* Text file busy */
	EFBIG   = 27 /* File too large */
	ENOSPC  = 28 /* No space left on device */
	ESPIPE  = 29 /* Illegal seek */
	EROFS   = 30 /* Read-only filesystem */
	EMLINK  = 31 /* Too many links */
	EPIPE   = 32 /* Broken pipe */

	/* math software */
	EDOM   = 33 /* Numerical argument out of domain */
	ERANGE = 34 /* Result too large */

	/* non-blocking and interrupt i/o */
	EAGAIN      = 35 /* Resource temporarily unavailable */
	EINPROGRESS = 36 /* Operation now in progress */
	EALREADY    = 37 /* Operation already in progress */

	/* ipc/network software -- argument errors */
	ENOTSOCK        = 38 /* Socket operation on non-socket */
	EDESTADDRREQ    = 39 /* Destination address required */
	EMSGSIZE        = 40 /* Message too long */
	EPROTOTYPE      = 41 /* Protocol wrong type for socket */
	ENOPROTOOPT     = 42 /* Protocol not available */
	EPROTONOSUPPORT = 43 /* Protocol not supported */
	ESOCKTNOSUPPORT = 44 /* Socket type not supported */
	EOPNOTSUPP      = 45 /* Operation not supported */
	EPFNOSUPPORT    = 46 /* Protocol family not supported */
	EAFNOSUPPORT    = 47 /* Address family not supported by protocol family */
	EADDRINUSE      = 48 /* Address already in use */
	EADDRNOTAVAIL   = 49 /* Can't assign requested address */

	/* ipc/network software -- operational errors */
	ENETDOWN     = 50 /* Network is down */
	ENETUNREACH  = 51 /* Network is unreachable */
	ENETRESET    = 52 /* Network dropped connection on reset */
	ECONNABORTED = 53 /* Software caused connection abort */
	ECONNRESET   = 54 /* Connection reset by peer */
	ENOBUFS      = 55 /* No buffer space available */
	EISCONN      = 56 /* Socket is already connected */
	ENOTCONN     = 57 /* Socket is not connected */
	ESHUTDOWN    = 58 /* Can't send after socket shutdown */
	ETOOMANYREFS = 59 /* Too many references: can't splice */
	ETIMEDOUT    = 60 /* Operation timed out */
	ECONNREFUSED = 61 /* Connection refused */

	ELOOP        = 62 /* Too many levels of symbolic links */
	ENAMETOOLONG = 63 /* File name too long */

	/* should be rearranged */
	EHOSTDOWN    = 64 /* Host is down */
	EHOSTUNREACH = 65 /* No route to host */
	ENOTEMPTY    = 66 /* Directory not empty */

	/* quotas & mush */
	EPROCLIM = 67 /* Too many processes */
	EUSERS   = 68 /* Too many users */
	EDQUOT   = 69 /* Disc quota exceeded */

	/* Network File System */
	ESTALE        = 70 /* Stale NFS file handle */
	EREMOTE       = 71 /* Too many levels of remote in path */
	EBADRPC       = 72 /* RPC struct is bad */
	ERPCMISMATCH  = 73 /* RPC version wrong */
	EPROGUNAVAIL  = 74 /* RPC prog. not avail */
	EPROGMISMATCH = 75 /* Program version wrong */
	EPROCUNAVAIL  = 76 /* Bad procedure for program */

	ENOLCK = 77 /* No locks available */
	ENOSYS = 78 /* Function not implemented */

	EFTYPE    = 79 /* Inappropriate file type or format */
	EAUTH     = 80 /* Authentication error */
	ENEEDAUTH = 81 /* Need authenticator */
	EIDRM     = 82 /* Identifier removed */
	ENOMSG    = 83 /* No message of desired type */
	EOVERFLOW = 84 /* Value too large to be stored in data type */
	ECANCELED = 85 /* Operation canceled */
	EILSEQ    = 86 /* Illegal byte sequence */
	ENOATTR   = 87 /* Attribute not found */

	EDOOFUS = 88 /* Programming error */

	EBADMSG   = 89 /* Bad message */
	EMULTIHOP = 90 /* Multihop attempted */
	ENOLINK   = 91 /* Link has been severed */
	EPROTO    = 92 /* Protocol error */

	ENOTCAPABLE     = 93 /* Capabilities insufficient */
	ECAPMODE        = 94 /* Not permitted in capability mode */
	ENOTRECOVERABLE = 95 /* State not recoverable */
	EOWNERDEAD      = 96 /* Previous owner died */
	EINTEGRITY      = 97 /* Integrity check failed */

	ELAST = 97 /* Must be equal largest errno */
)

var errorStrings = [...]string{
	"",
	EPERM:   " Operation not permitted ",
	ENOENT:  " No such file or directory ",
	ESRCH:   " No such process ",
	EINTR:   " Interrupted system call ",
	EIO:     " Input/output error ",
	ENXIO:   " Device not configured ",
	E2BIG:   " Argument list too long ",
	ENOEXEC: " Exec format error ",
	EBADF:   " Bad file descriptor ",
	ECHILD:  " No child processes ",
	EDEADLK: " Resource deadlock avoided ",
	/* 11 was EAGAIN */
	ENOMEM:  " Cannot allocate memory ",
	EACCES:  " Permission denied ",
	EFAULT:  " Bad address ",
	ENOTBLK: " Block device required ",
	EBUSY:   " Device busy ",
	EEXIST:  " File exists ",
	EXDEV:   " Cross-device link ",
	ENODEV:  " Operation not supported by device ",
	ENOTDIR: " Not a directory ",
	EISDIR:  " Is a directory ",
	EINVAL:  " Invalid argument ",
	ENFILE:  " Too many open files in system ",
	EMFILE:  " Too many open files ",
	ENOTTY:  " Inappropriate ioctl for device ",
	ETXTBSY: " Text file busy ",
	EFBIG:   " File too large ",
	ENOSPC:  " No space left on device ",
	ESPIPE:  " Illegal seek ",
	EROFS:   " Read-only filesystem ",
	EMLINK:  " Too many links ",
	EPIPE:   " Broken pipe ",

	/* math software */
	EDOM:   " Numerical argument out of domain ",
	ERANGE: " Result too large ",

	/* non-blocking and interrupt i/o */
	EAGAIN:      " Resource temporarily unavailable ",
	EINPROGRESS: "Operation now in progress",
	EALREADY:    "Operation already in progress",

	/* ipc/network software -- argument errors */
	ENOTSOCK:        " Socket operation on non-socket ",
	EDESTADDRREQ:    " Destination address required ",
	EMSGSIZE:        " Message too long ",
	EPROTOTYPE:      " Protocol wrong type for socket ",
	ENOPROTOOPT:     " Protocol not available ",
	EPROTONOSUPPORT: " Protocol not supported ",
	ESOCKTNOSUPPORT: " Socket type not supported ",
	EOPNOTSUPP:      " Operation not supported ",
	EPFNOSUPPORT:    " Protocol family not supported ",
	EAFNOSUPPORT:    " Address family not supported by protocol family ",
	EADDRINUSE:      " Address already in use ",
	EADDRNOTAVAIL:   " Can't assign requested address ",

	/* ipc/network software -- operational errors */
	ENETDOWN:     " Network is down ",
	ENETUNREACH:  " Network is unreachable ",
	ENETRESET:    " Network dropped connection on reset ",
	ECONNABORTED: " Software caused connection abort ",
	ECONNRESET:   " Connection reset by peer ",
	ENOBUFS:      " No buffer space available ",
	EISCONN:      " Socket is already connected ",
	ENOTCONN:     " Socket is not connected ",
	ESHUTDOWN:    " Can't send after socket shutdown ",
	ETOOMANYREFS: " Too many references: can't splice ",
	ETIMEDOUT:    " Operation timed out ",
	ECONNREFUSED: " Connection refused ",

	ELOOP:        " Too many levels of symbolic links ",
	ENAMETOOLONG: " File name too long ",

	/* should be rearranged */
	EHOSTDOWN:    " Host is down ",
	EHOSTUNREACH: " No route to host ",
	ENOTEMPTY:    " Directory not empty ",

	/* quotas & mush */
	EPROCLIM: " Too many processes ",
	EUSERS:   " Too many users ",
	EDQUOT:   " Disc quota exceeded ",

	/* Network File System */
	ESTALE:        " Stale NFS file handle ",
	EREMOTE:       " Too many levels of remote in path ",
	EBADRPC:       " RPC struct is bad ",
	ERPCMISMATCH:  " RPC version wrong ",
	EPROGUNAVAIL:  " RPC prog. not avail ",
	EPROGMISMATCH: " Program version wrong ",
	EPROCUNAVAIL:  " Bad procedure for program ",

	ENOLCK: " No locks available ",
	ENOSYS: " Function not implemented ",

	EFTYPE:    " Inappropriate file type or format ",
	EAUTH:     " Authentication error ",
	ENEEDAUTH: " Need authenticator ",
	EIDRM:     " Identifier removed ",
	ENOMSG:    " No message of desired type ",
	EOVERFLOW: " Value too large to be stored in data type ",
	ECANCELED: " Operation canceled ",
	EILSEQ:    " Illegal byte sequence ",
	ENOATTR:   " Attribute not found ",

	EDOOFUS: " Programming error ",

	EBADMSG:   " Bad message ",
	EMULTIHOP: " Multihop attempted ",
	ENOLINK:   " Link has been severed ",
	EPROTO:    " Protocol error ",

	ENOTCAPABLE:     " Capabilities insufficient ",
	ECAPMODE:        " Not permitted in capability mode ",
	ENOTRECOVERABLE: " State not recoverable ",
	EOWNERDEAD:      " Previous owner died ",
	EINTEGRITY:      " Integrity check failed ",
}

func Strerror(errno int) string {
	if (errno >= 0) && (errno <= ELAST) {
		return errorStrings[errno]
	} else {
		return "<UNDEFINED ERROR>"
	}
}
