// Automatically generated by the res2go IDE plug-in, do not edit.
package main

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TForm1 struct {
    *vcl.TForm
    Button1 *vcl.TButton

    //::private::
    TForm1Fields
}

var Form1 *TForm1




// vcl.Application.CreateForm(&Form1)

func NewForm1(owner vcl.IComponent) (root *TForm1)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/unit1.gfm
var form1Bytes []byte

// Register Form Resources
var _ = vcl.RegisterFormResource(Form1, &form1Bytes)
