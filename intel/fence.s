#include "textflag.h"

TEXT ·LFENCE(SB), NOSPLIT, $0
	LFENCE
	RET

TEXT ·MFENCE(SB), NOSPLIT, $0
	MFENCE
	RET

TEXT ·SFENCE(SB), NOSPLIT, $0
	SFENCE
	RET
