//	Auto-generated by the "go-xsd" package located at:
//		github.com/metaleap/go-xsd
//	Comments on types and fields (if any) are from the XSD file located at:
//		www.w3.org/Math/XMLSchema/mathml2/content/linear-algebra.xsd
package gopkg_WwwW3OrgMathXMLSchemaMathml2Mathml2Xsd

//	This is an XML Schema module for the linear algebra part of content MathML.
//	Author: Stéphane Dalmas, INRIA.
import (
)

type XsdGoPkgHasAtts_VectorAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_MatrixAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_MatrixrowAttlist struct {
	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_DeterminantAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_TransposeAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_SelectorAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_VectorproductAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_ScalarproductAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type XsdGoPkgHasAtts_OuterproductAttlist struct {
	XsdGoPkgHasAtts_DefinitionAttrib

	XsdGoPkgHasAtts_CommonAttrib

}

type TvectorType struct {
	XsdGoPkgHasGroup_VectorContent

	XsdGoPkgHasAtts_VectorAttlist

}

type TmatrixType struct {
	XsdGoPkgHasGroup_MatrixContent

	XsdGoPkgHasAtts_MatrixAttlist

}

type TmatrixrowType struct {
	XsdGoPkgHasGroup_MatrixrowContent

	XsdGoPkgHasAtts_MatrixrowAttlist

}

type TdeterminantType struct {
	XsdGoPkgHasAtts_DeterminantAttlist

}

type TtransposeType struct {
	XsdGoPkgHasAtts_TransposeAttlist

}

type TselectorType struct {
	XsdGoPkgHasAtts_SelectorAttlist

}

type TvectorproductType struct {
	XsdGoPkgHasAtts_VectorproductAttlist

}

type TscalarproductType struct {
	XsdGoPkgHasAtts_ScalarproductAttlist

}

type TouterproductType struct {
	XsdGoPkgHasAtts_OuterproductAttlist

}

type XsdGoPkgHasElem_Vector struct {
	Vector *TvectorType `xml:"http://www.w3.org/1998/Math/MathML vector"`
}

type XsdGoPkgHasElems_Vector struct {
	Vectors []*TvectorType `xml:"http://www.w3.org/1998/Math/MathML vector"`
}

type XsdGoPkgHasElem_Matrix struct {
	Matrix *TmatrixType `xml:"http://www.w3.org/1998/Math/MathML matrix"`
}

type XsdGoPkgHasElems_Matrix struct {
	Matrixs []*TmatrixType `xml:"http://www.w3.org/1998/Math/MathML matrix"`
}

type XsdGoPkgHasElems_Matrixrow struct {
	Matrixrows []*TmatrixrowType `xml:"http://www.w3.org/1998/Math/MathML matrixrow"`
}

type XsdGoPkgHasElem_Matrixrow struct {
	Matrixrow *TmatrixrowType `xml:"http://www.w3.org/1998/Math/MathML matrixrow"`
}

type XsdGoPkgHasElems_Determinant struct {
	Determinants []*TdeterminantType `xml:"http://www.w3.org/1998/Math/MathML determinant"`
}

type XsdGoPkgHasElem_Determinant struct {
	Determinant *TdeterminantType `xml:"http://www.w3.org/1998/Math/MathML determinant"`
}

type XsdGoPkgHasElems_Transpose struct {
	Transposes []*TtransposeType `xml:"http://www.w3.org/1998/Math/MathML transpose"`
}

type XsdGoPkgHasElem_Transpose struct {
	Transpose *TtransposeType `xml:"http://www.w3.org/1998/Math/MathML transpose"`
}

type XsdGoPkgHasElem_Selector struct {
	Selector *TselectorType `xml:"http://www.w3.org/1998/Math/MathML selector"`
}

type XsdGoPkgHasElems_Selector struct {
	Selectors []*TselectorType `xml:"http://www.w3.org/1998/Math/MathML selector"`
}

type XsdGoPkgHasElem_Vectorproduct struct {
	Vectorproduct *TvectorproductType `xml:"http://www.w3.org/1998/Math/MathML vectorproduct"`
}

type XsdGoPkgHasElems_Vectorproduct struct {
	Vectorproducts []*TvectorproductType `xml:"http://www.w3.org/1998/Math/MathML vectorproduct"`
}

type XsdGoPkgHasElem_Scalarproduct struct {
	Scalarproduct *TscalarproductType `xml:"http://www.w3.org/1998/Math/MathML scalarproduct"`
}

type XsdGoPkgHasElems_Scalarproduct struct {
	Scalarproducts []*TscalarproductType `xml:"http://www.w3.org/1998/Math/MathML scalarproduct"`
}

type XsdGoPkgHasElems_Outerproduct struct {
	Outerproducts []*TouterproductType `xml:"http://www.w3.org/1998/Math/MathML outerproduct"`
}

type XsdGoPkgHasElem_Outerproduct struct {
	Outerproduct *TouterproductType `xml:"http://www.w3.org/1998/Math/MathML outerproduct"`
}

type XsdGoPkgHasGroup_VectorContent struct {
	XsdGoPkgHasGroup_ContentExprClass

}

type XsdGoPkgHasGroup_MatrixContent struct {
	XsdGoPkgHasElem_Matrixrow

}

type XsdGoPkgHasGroup_MatrixrowContent struct {
	XsdGoPkgHasGroup_ContentExprClass

}

type XsdGoPkgHasGroup_ContentLinearAlgebraClass struct {
	XsdGoPkgHasElem_Vector

	XsdGoPkgHasElem_Matrix

	XsdGoPkgHasElem_Determinant

	XsdGoPkgHasElem_Transpose

	XsdGoPkgHasElem_Selector

	XsdGoPkgHasElem_Vectorproduct

	XsdGoPkgHasElem_Scalarproduct

	XsdGoPkgHasElem_Outerproduct

}
