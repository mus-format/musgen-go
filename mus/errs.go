package musgen

import "fmt"

var ErrCodeGenFailed = fmt.Errorf("code generation failed: ensure all musgen.CodeGenerator " +
	"options are set correctly; see the generated code for details")
