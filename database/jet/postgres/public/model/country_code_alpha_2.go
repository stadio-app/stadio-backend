//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import "errors"

type CountryCodeAlpha2 string

const (
	CountryCodeAlpha2_Af CountryCodeAlpha2 = "AF"
	CountryCodeAlpha2_Ax CountryCodeAlpha2 = "AX"
	CountryCodeAlpha2_Al CountryCodeAlpha2 = "AL"
	CountryCodeAlpha2_Dz CountryCodeAlpha2 = "DZ"
	CountryCodeAlpha2_As CountryCodeAlpha2 = "AS"
	CountryCodeAlpha2_Ad CountryCodeAlpha2 = "AD"
	CountryCodeAlpha2_Ao CountryCodeAlpha2 = "AO"
	CountryCodeAlpha2_Ai CountryCodeAlpha2 = "AI"
	CountryCodeAlpha2_Aq CountryCodeAlpha2 = "AQ"
	CountryCodeAlpha2_Ag CountryCodeAlpha2 = "AG"
	CountryCodeAlpha2_Ar CountryCodeAlpha2 = "AR"
	CountryCodeAlpha2_Am CountryCodeAlpha2 = "AM"
	CountryCodeAlpha2_Aw CountryCodeAlpha2 = "AW"
	CountryCodeAlpha2_Au CountryCodeAlpha2 = "AU"
	CountryCodeAlpha2_At CountryCodeAlpha2 = "AT"
	CountryCodeAlpha2_Az CountryCodeAlpha2 = "AZ"
	CountryCodeAlpha2_Bs CountryCodeAlpha2 = "BS"
	CountryCodeAlpha2_Bh CountryCodeAlpha2 = "BH"
	CountryCodeAlpha2_Bd CountryCodeAlpha2 = "BD"
	CountryCodeAlpha2_Bb CountryCodeAlpha2 = "BB"
	CountryCodeAlpha2_By CountryCodeAlpha2 = "BY"
	CountryCodeAlpha2_Be CountryCodeAlpha2 = "BE"
	CountryCodeAlpha2_Bz CountryCodeAlpha2 = "BZ"
	CountryCodeAlpha2_Bj CountryCodeAlpha2 = "BJ"
	CountryCodeAlpha2_Bm CountryCodeAlpha2 = "BM"
	CountryCodeAlpha2_Bt CountryCodeAlpha2 = "BT"
	CountryCodeAlpha2_Bo CountryCodeAlpha2 = "BO"
	CountryCodeAlpha2_Ba CountryCodeAlpha2 = "BA"
	CountryCodeAlpha2_Bw CountryCodeAlpha2 = "BW"
	CountryCodeAlpha2_Bv CountryCodeAlpha2 = "BV"
	CountryCodeAlpha2_Br CountryCodeAlpha2 = "BR"
	CountryCodeAlpha2_Io CountryCodeAlpha2 = "IO"
	CountryCodeAlpha2_Bn CountryCodeAlpha2 = "BN"
	CountryCodeAlpha2_Bg CountryCodeAlpha2 = "BG"
	CountryCodeAlpha2_Bf CountryCodeAlpha2 = "BF"
	CountryCodeAlpha2_Bi CountryCodeAlpha2 = "BI"
	CountryCodeAlpha2_Kh CountryCodeAlpha2 = "KH"
	CountryCodeAlpha2_Cm CountryCodeAlpha2 = "CM"
	CountryCodeAlpha2_Ca CountryCodeAlpha2 = "CA"
	CountryCodeAlpha2_Cv CountryCodeAlpha2 = "CV"
	CountryCodeAlpha2_Ky CountryCodeAlpha2 = "KY"
	CountryCodeAlpha2_Cf CountryCodeAlpha2 = "CF"
	CountryCodeAlpha2_Td CountryCodeAlpha2 = "TD"
	CountryCodeAlpha2_Cl CountryCodeAlpha2 = "CL"
	CountryCodeAlpha2_Cn CountryCodeAlpha2 = "CN"
	CountryCodeAlpha2_Cx CountryCodeAlpha2 = "CX"
	CountryCodeAlpha2_Cc CountryCodeAlpha2 = "CC"
	CountryCodeAlpha2_Co CountryCodeAlpha2 = "CO"
	CountryCodeAlpha2_Km CountryCodeAlpha2 = "KM"
	CountryCodeAlpha2_Cg CountryCodeAlpha2 = "CG"
	CountryCodeAlpha2_Cd CountryCodeAlpha2 = "CD"
	CountryCodeAlpha2_Ck CountryCodeAlpha2 = "CK"
	CountryCodeAlpha2_Cr CountryCodeAlpha2 = "CR"
	CountryCodeAlpha2_Ci CountryCodeAlpha2 = "CI"
	CountryCodeAlpha2_Hr CountryCodeAlpha2 = "HR"
	CountryCodeAlpha2_Cu CountryCodeAlpha2 = "CU"
	CountryCodeAlpha2_Cy CountryCodeAlpha2 = "CY"
	CountryCodeAlpha2_Cz CountryCodeAlpha2 = "CZ"
	CountryCodeAlpha2_Dk CountryCodeAlpha2 = "DK"
	CountryCodeAlpha2_Dj CountryCodeAlpha2 = "DJ"
	CountryCodeAlpha2_Dm CountryCodeAlpha2 = "DM"
	CountryCodeAlpha2_Do CountryCodeAlpha2 = "DO"
	CountryCodeAlpha2_Ec CountryCodeAlpha2 = "EC"
	CountryCodeAlpha2_Eg CountryCodeAlpha2 = "EG"
	CountryCodeAlpha2_Sv CountryCodeAlpha2 = "SV"
	CountryCodeAlpha2_Gq CountryCodeAlpha2 = "GQ"
	CountryCodeAlpha2_Er CountryCodeAlpha2 = "ER"
	CountryCodeAlpha2_Ee CountryCodeAlpha2 = "EE"
	CountryCodeAlpha2_Et CountryCodeAlpha2 = "ET"
	CountryCodeAlpha2_Fk CountryCodeAlpha2 = "FK"
	CountryCodeAlpha2_Fo CountryCodeAlpha2 = "FO"
	CountryCodeAlpha2_Fj CountryCodeAlpha2 = "FJ"
	CountryCodeAlpha2_Fi CountryCodeAlpha2 = "FI"
	CountryCodeAlpha2_Fr CountryCodeAlpha2 = "FR"
	CountryCodeAlpha2_Gf CountryCodeAlpha2 = "GF"
	CountryCodeAlpha2_Pf CountryCodeAlpha2 = "PF"
	CountryCodeAlpha2_Tf CountryCodeAlpha2 = "TF"
	CountryCodeAlpha2_Ga CountryCodeAlpha2 = "GA"
	CountryCodeAlpha2_Gm CountryCodeAlpha2 = "GM"
	CountryCodeAlpha2_Ge CountryCodeAlpha2 = "GE"
	CountryCodeAlpha2_De CountryCodeAlpha2 = "DE"
	CountryCodeAlpha2_Gh CountryCodeAlpha2 = "GH"
	CountryCodeAlpha2_Gi CountryCodeAlpha2 = "GI"
	CountryCodeAlpha2_Gr CountryCodeAlpha2 = "GR"
	CountryCodeAlpha2_Gl CountryCodeAlpha2 = "GL"
	CountryCodeAlpha2_Gd CountryCodeAlpha2 = "GD"
	CountryCodeAlpha2_Gp CountryCodeAlpha2 = "GP"
	CountryCodeAlpha2_Gu CountryCodeAlpha2 = "GU"
	CountryCodeAlpha2_Gt CountryCodeAlpha2 = "GT"
	CountryCodeAlpha2_Gg CountryCodeAlpha2 = "GG"
	CountryCodeAlpha2_Gn CountryCodeAlpha2 = "GN"
	CountryCodeAlpha2_Gw CountryCodeAlpha2 = "GW"
	CountryCodeAlpha2_Gy CountryCodeAlpha2 = "GY"
	CountryCodeAlpha2_Ht CountryCodeAlpha2 = "HT"
	CountryCodeAlpha2_Hm CountryCodeAlpha2 = "HM"
	CountryCodeAlpha2_Va CountryCodeAlpha2 = "VA"
	CountryCodeAlpha2_Hn CountryCodeAlpha2 = "HN"
	CountryCodeAlpha2_Hk CountryCodeAlpha2 = "HK"
	CountryCodeAlpha2_Hu CountryCodeAlpha2 = "HU"
	CountryCodeAlpha2_Is CountryCodeAlpha2 = "IS"
	CountryCodeAlpha2_In CountryCodeAlpha2 = "IN"
	CountryCodeAlpha2_ID CountryCodeAlpha2 = "ID"
	CountryCodeAlpha2_Ir CountryCodeAlpha2 = "IR"
	CountryCodeAlpha2_Iq CountryCodeAlpha2 = "IQ"
	CountryCodeAlpha2_Ie CountryCodeAlpha2 = "IE"
	CountryCodeAlpha2_Im CountryCodeAlpha2 = "IM"
	CountryCodeAlpha2_Il CountryCodeAlpha2 = "IL"
	CountryCodeAlpha2_It CountryCodeAlpha2 = "IT"
	CountryCodeAlpha2_Jm CountryCodeAlpha2 = "JM"
	CountryCodeAlpha2_Jp CountryCodeAlpha2 = "JP"
	CountryCodeAlpha2_Je CountryCodeAlpha2 = "JE"
	CountryCodeAlpha2_Jo CountryCodeAlpha2 = "JO"
	CountryCodeAlpha2_Kz CountryCodeAlpha2 = "KZ"
	CountryCodeAlpha2_Ke CountryCodeAlpha2 = "KE"
	CountryCodeAlpha2_Ki CountryCodeAlpha2 = "KI"
	CountryCodeAlpha2_Kr CountryCodeAlpha2 = "KR"
	CountryCodeAlpha2_Kp CountryCodeAlpha2 = "KP"
	CountryCodeAlpha2_Kw CountryCodeAlpha2 = "KW"
	CountryCodeAlpha2_Kg CountryCodeAlpha2 = "KG"
	CountryCodeAlpha2_La CountryCodeAlpha2 = "LA"
	CountryCodeAlpha2_Lv CountryCodeAlpha2 = "LV"
	CountryCodeAlpha2_Lb CountryCodeAlpha2 = "LB"
	CountryCodeAlpha2_Ls CountryCodeAlpha2 = "LS"
	CountryCodeAlpha2_Lr CountryCodeAlpha2 = "LR"
	CountryCodeAlpha2_Ly CountryCodeAlpha2 = "LY"
	CountryCodeAlpha2_Li CountryCodeAlpha2 = "LI"
	CountryCodeAlpha2_Lt CountryCodeAlpha2 = "LT"
	CountryCodeAlpha2_Lu CountryCodeAlpha2 = "LU"
	CountryCodeAlpha2_Mo CountryCodeAlpha2 = "MO"
	CountryCodeAlpha2_Mk CountryCodeAlpha2 = "MK"
	CountryCodeAlpha2_Mg CountryCodeAlpha2 = "MG"
	CountryCodeAlpha2_Mw CountryCodeAlpha2 = "MW"
	CountryCodeAlpha2_My CountryCodeAlpha2 = "MY"
	CountryCodeAlpha2_Mv CountryCodeAlpha2 = "MV"
	CountryCodeAlpha2_Ml CountryCodeAlpha2 = "ML"
	CountryCodeAlpha2_Mt CountryCodeAlpha2 = "MT"
	CountryCodeAlpha2_Mh CountryCodeAlpha2 = "MH"
	CountryCodeAlpha2_Mq CountryCodeAlpha2 = "MQ"
	CountryCodeAlpha2_Mr CountryCodeAlpha2 = "MR"
	CountryCodeAlpha2_Mu CountryCodeAlpha2 = "MU"
	CountryCodeAlpha2_Yt CountryCodeAlpha2 = "YT"
	CountryCodeAlpha2_Mx CountryCodeAlpha2 = "MX"
	CountryCodeAlpha2_Fm CountryCodeAlpha2 = "FM"
	CountryCodeAlpha2_Md CountryCodeAlpha2 = "MD"
	CountryCodeAlpha2_Mc CountryCodeAlpha2 = "MC"
	CountryCodeAlpha2_Mn CountryCodeAlpha2 = "MN"
	CountryCodeAlpha2_Me CountryCodeAlpha2 = "ME"
	CountryCodeAlpha2_Ms CountryCodeAlpha2 = "MS"
	CountryCodeAlpha2_Ma CountryCodeAlpha2 = "MA"
	CountryCodeAlpha2_Mz CountryCodeAlpha2 = "MZ"
	CountryCodeAlpha2_Mm CountryCodeAlpha2 = "MM"
	CountryCodeAlpha2_Na CountryCodeAlpha2 = "NA"
	CountryCodeAlpha2_Nr CountryCodeAlpha2 = "NR"
	CountryCodeAlpha2_Np CountryCodeAlpha2 = "NP"
	CountryCodeAlpha2_Nl CountryCodeAlpha2 = "NL"
	CountryCodeAlpha2_An CountryCodeAlpha2 = "AN"
	CountryCodeAlpha2_Nc CountryCodeAlpha2 = "NC"
	CountryCodeAlpha2_Nz CountryCodeAlpha2 = "NZ"
	CountryCodeAlpha2_Ni CountryCodeAlpha2 = "NI"
	CountryCodeAlpha2_Ne CountryCodeAlpha2 = "NE"
	CountryCodeAlpha2_Ng CountryCodeAlpha2 = "NG"
	CountryCodeAlpha2_Nu CountryCodeAlpha2 = "NU"
	CountryCodeAlpha2_Nf CountryCodeAlpha2 = "NF"
	CountryCodeAlpha2_Mp CountryCodeAlpha2 = "MP"
	CountryCodeAlpha2_No CountryCodeAlpha2 = "NO"
	CountryCodeAlpha2_Om CountryCodeAlpha2 = "OM"
	CountryCodeAlpha2_Pk CountryCodeAlpha2 = "PK"
	CountryCodeAlpha2_Pw CountryCodeAlpha2 = "PW"
	CountryCodeAlpha2_Ps CountryCodeAlpha2 = "PS"
	CountryCodeAlpha2_Pa CountryCodeAlpha2 = "PA"
	CountryCodeAlpha2_Pg CountryCodeAlpha2 = "PG"
	CountryCodeAlpha2_Py CountryCodeAlpha2 = "PY"
	CountryCodeAlpha2_Pe CountryCodeAlpha2 = "PE"
	CountryCodeAlpha2_Ph CountryCodeAlpha2 = "PH"
	CountryCodeAlpha2_Pn CountryCodeAlpha2 = "PN"
	CountryCodeAlpha2_Pl CountryCodeAlpha2 = "PL"
	CountryCodeAlpha2_Pt CountryCodeAlpha2 = "PT"
	CountryCodeAlpha2_Pr CountryCodeAlpha2 = "PR"
	CountryCodeAlpha2_Qa CountryCodeAlpha2 = "QA"
	CountryCodeAlpha2_Re CountryCodeAlpha2 = "RE"
	CountryCodeAlpha2_Ro CountryCodeAlpha2 = "RO"
	CountryCodeAlpha2_Ru CountryCodeAlpha2 = "RU"
	CountryCodeAlpha2_Rw CountryCodeAlpha2 = "RW"
	CountryCodeAlpha2_Bl CountryCodeAlpha2 = "BL"
	CountryCodeAlpha2_Sh CountryCodeAlpha2 = "SH"
	CountryCodeAlpha2_Kn CountryCodeAlpha2 = "KN"
	CountryCodeAlpha2_Lc CountryCodeAlpha2 = "LC"
	CountryCodeAlpha2_Mf CountryCodeAlpha2 = "MF"
	CountryCodeAlpha2_Pm CountryCodeAlpha2 = "PM"
	CountryCodeAlpha2_Vc CountryCodeAlpha2 = "VC"
	CountryCodeAlpha2_Ws CountryCodeAlpha2 = "WS"
	CountryCodeAlpha2_Sm CountryCodeAlpha2 = "SM"
	CountryCodeAlpha2_St CountryCodeAlpha2 = "ST"
	CountryCodeAlpha2_Sa CountryCodeAlpha2 = "SA"
	CountryCodeAlpha2_Sn CountryCodeAlpha2 = "SN"
	CountryCodeAlpha2_Rs CountryCodeAlpha2 = "RS"
	CountryCodeAlpha2_Sc CountryCodeAlpha2 = "SC"
	CountryCodeAlpha2_Sl CountryCodeAlpha2 = "SL"
	CountryCodeAlpha2_Sg CountryCodeAlpha2 = "SG"
	CountryCodeAlpha2_Sk CountryCodeAlpha2 = "SK"
	CountryCodeAlpha2_Si CountryCodeAlpha2 = "SI"
	CountryCodeAlpha2_Sb CountryCodeAlpha2 = "SB"
	CountryCodeAlpha2_So CountryCodeAlpha2 = "SO"
	CountryCodeAlpha2_Za CountryCodeAlpha2 = "ZA"
	CountryCodeAlpha2_Gs CountryCodeAlpha2 = "GS"
	CountryCodeAlpha2_Es CountryCodeAlpha2 = "ES"
	CountryCodeAlpha2_Lk CountryCodeAlpha2 = "LK"
	CountryCodeAlpha2_Sd CountryCodeAlpha2 = "SD"
	CountryCodeAlpha2_Sr CountryCodeAlpha2 = "SR"
	CountryCodeAlpha2_Sj CountryCodeAlpha2 = "SJ"
	CountryCodeAlpha2_Sz CountryCodeAlpha2 = "SZ"
	CountryCodeAlpha2_Se CountryCodeAlpha2 = "SE"
	CountryCodeAlpha2_Ch CountryCodeAlpha2 = "CH"
	CountryCodeAlpha2_Sy CountryCodeAlpha2 = "SY"
	CountryCodeAlpha2_Tw CountryCodeAlpha2 = "TW"
	CountryCodeAlpha2_Tj CountryCodeAlpha2 = "TJ"
	CountryCodeAlpha2_Tz CountryCodeAlpha2 = "TZ"
	CountryCodeAlpha2_Th CountryCodeAlpha2 = "TH"
	CountryCodeAlpha2_Tl CountryCodeAlpha2 = "TL"
	CountryCodeAlpha2_Tg CountryCodeAlpha2 = "TG"
	CountryCodeAlpha2_Tk CountryCodeAlpha2 = "TK"
	CountryCodeAlpha2_To CountryCodeAlpha2 = "TO"
	CountryCodeAlpha2_Tt CountryCodeAlpha2 = "TT"
	CountryCodeAlpha2_Tn CountryCodeAlpha2 = "TN"
	CountryCodeAlpha2_Tr CountryCodeAlpha2 = "TR"
	CountryCodeAlpha2_Tm CountryCodeAlpha2 = "TM"
	CountryCodeAlpha2_Tc CountryCodeAlpha2 = "TC"
	CountryCodeAlpha2_Tv CountryCodeAlpha2 = "TV"
	CountryCodeAlpha2_Ug CountryCodeAlpha2 = "UG"
	CountryCodeAlpha2_Ua CountryCodeAlpha2 = "UA"
	CountryCodeAlpha2_Ae CountryCodeAlpha2 = "AE"
	CountryCodeAlpha2_Gb CountryCodeAlpha2 = "GB"
	CountryCodeAlpha2_Us CountryCodeAlpha2 = "US"
	CountryCodeAlpha2_Um CountryCodeAlpha2 = "UM"
	CountryCodeAlpha2_Uy CountryCodeAlpha2 = "UY"
	CountryCodeAlpha2_Uz CountryCodeAlpha2 = "UZ"
	CountryCodeAlpha2_Vu CountryCodeAlpha2 = "VU"
	CountryCodeAlpha2_Ve CountryCodeAlpha2 = "VE"
	CountryCodeAlpha2_Vn CountryCodeAlpha2 = "VN"
	CountryCodeAlpha2_Vg CountryCodeAlpha2 = "VG"
	CountryCodeAlpha2_Vi CountryCodeAlpha2 = "VI"
	CountryCodeAlpha2_Wf CountryCodeAlpha2 = "WF"
	CountryCodeAlpha2_Eh CountryCodeAlpha2 = "EH"
	CountryCodeAlpha2_Ye CountryCodeAlpha2 = "YE"
	CountryCodeAlpha2_Zm CountryCodeAlpha2 = "ZM"
	CountryCodeAlpha2_Zw CountryCodeAlpha2 = "ZW"
	CountryCodeAlpha2_Ss CountryCodeAlpha2 = "SS"
	CountryCodeAlpha2_Xk CountryCodeAlpha2 = "XK"
	CountryCodeAlpha2_Bq CountryCodeAlpha2 = "BQ"
)

func (e *CountryCodeAlpha2) Scan(value interface{}) error {
	var enumValue string
	switch val := value.(type) {
	case string:
		enumValue = val
	case []byte:
		enumValue = string(val)
	default:
		return errors.New("jet: Invalid scan value for AllTypesEnum enum. Enum value has to be of type string or []byte")
	}

	switch enumValue {
	case "AF":
		*e = CountryCodeAlpha2_Af
	case "AX":
		*e = CountryCodeAlpha2_Ax
	case "AL":
		*e = CountryCodeAlpha2_Al
	case "DZ":
		*e = CountryCodeAlpha2_Dz
	case "AS":
		*e = CountryCodeAlpha2_As
	case "AD":
		*e = CountryCodeAlpha2_Ad
	case "AO":
		*e = CountryCodeAlpha2_Ao
	case "AI":
		*e = CountryCodeAlpha2_Ai
	case "AQ":
		*e = CountryCodeAlpha2_Aq
	case "AG":
		*e = CountryCodeAlpha2_Ag
	case "AR":
		*e = CountryCodeAlpha2_Ar
	case "AM":
		*e = CountryCodeAlpha2_Am
	case "AW":
		*e = CountryCodeAlpha2_Aw
	case "AU":
		*e = CountryCodeAlpha2_Au
	case "AT":
		*e = CountryCodeAlpha2_At
	case "AZ":
		*e = CountryCodeAlpha2_Az
	case "BS":
		*e = CountryCodeAlpha2_Bs
	case "BH":
		*e = CountryCodeAlpha2_Bh
	case "BD":
		*e = CountryCodeAlpha2_Bd
	case "BB":
		*e = CountryCodeAlpha2_Bb
	case "BY":
		*e = CountryCodeAlpha2_By
	case "BE":
		*e = CountryCodeAlpha2_Be
	case "BZ":
		*e = CountryCodeAlpha2_Bz
	case "BJ":
		*e = CountryCodeAlpha2_Bj
	case "BM":
		*e = CountryCodeAlpha2_Bm
	case "BT":
		*e = CountryCodeAlpha2_Bt
	case "BO":
		*e = CountryCodeAlpha2_Bo
	case "BA":
		*e = CountryCodeAlpha2_Ba
	case "BW":
		*e = CountryCodeAlpha2_Bw
	case "BV":
		*e = CountryCodeAlpha2_Bv
	case "BR":
		*e = CountryCodeAlpha2_Br
	case "IO":
		*e = CountryCodeAlpha2_Io
	case "BN":
		*e = CountryCodeAlpha2_Bn
	case "BG":
		*e = CountryCodeAlpha2_Bg
	case "BF":
		*e = CountryCodeAlpha2_Bf
	case "BI":
		*e = CountryCodeAlpha2_Bi
	case "KH":
		*e = CountryCodeAlpha2_Kh
	case "CM":
		*e = CountryCodeAlpha2_Cm
	case "CA":
		*e = CountryCodeAlpha2_Ca
	case "CV":
		*e = CountryCodeAlpha2_Cv
	case "KY":
		*e = CountryCodeAlpha2_Ky
	case "CF":
		*e = CountryCodeAlpha2_Cf
	case "TD":
		*e = CountryCodeAlpha2_Td
	case "CL":
		*e = CountryCodeAlpha2_Cl
	case "CN":
		*e = CountryCodeAlpha2_Cn
	case "CX":
		*e = CountryCodeAlpha2_Cx
	case "CC":
		*e = CountryCodeAlpha2_Cc
	case "CO":
		*e = CountryCodeAlpha2_Co
	case "KM":
		*e = CountryCodeAlpha2_Km
	case "CG":
		*e = CountryCodeAlpha2_Cg
	case "CD":
		*e = CountryCodeAlpha2_Cd
	case "CK":
		*e = CountryCodeAlpha2_Ck
	case "CR":
		*e = CountryCodeAlpha2_Cr
	case "CI":
		*e = CountryCodeAlpha2_Ci
	case "HR":
		*e = CountryCodeAlpha2_Hr
	case "CU":
		*e = CountryCodeAlpha2_Cu
	case "CY":
		*e = CountryCodeAlpha2_Cy
	case "CZ":
		*e = CountryCodeAlpha2_Cz
	case "DK":
		*e = CountryCodeAlpha2_Dk
	case "DJ":
		*e = CountryCodeAlpha2_Dj
	case "DM":
		*e = CountryCodeAlpha2_Dm
	case "DO":
		*e = CountryCodeAlpha2_Do
	case "EC":
		*e = CountryCodeAlpha2_Ec
	case "EG":
		*e = CountryCodeAlpha2_Eg
	case "SV":
		*e = CountryCodeAlpha2_Sv
	case "GQ":
		*e = CountryCodeAlpha2_Gq
	case "ER":
		*e = CountryCodeAlpha2_Er
	case "EE":
		*e = CountryCodeAlpha2_Ee
	case "ET":
		*e = CountryCodeAlpha2_Et
	case "FK":
		*e = CountryCodeAlpha2_Fk
	case "FO":
		*e = CountryCodeAlpha2_Fo
	case "FJ":
		*e = CountryCodeAlpha2_Fj
	case "FI":
		*e = CountryCodeAlpha2_Fi
	case "FR":
		*e = CountryCodeAlpha2_Fr
	case "GF":
		*e = CountryCodeAlpha2_Gf
	case "PF":
		*e = CountryCodeAlpha2_Pf
	case "TF":
		*e = CountryCodeAlpha2_Tf
	case "GA":
		*e = CountryCodeAlpha2_Ga
	case "GM":
		*e = CountryCodeAlpha2_Gm
	case "GE":
		*e = CountryCodeAlpha2_Ge
	case "DE":
		*e = CountryCodeAlpha2_De
	case "GH":
		*e = CountryCodeAlpha2_Gh
	case "GI":
		*e = CountryCodeAlpha2_Gi
	case "GR":
		*e = CountryCodeAlpha2_Gr
	case "GL":
		*e = CountryCodeAlpha2_Gl
	case "GD":
		*e = CountryCodeAlpha2_Gd
	case "GP":
		*e = CountryCodeAlpha2_Gp
	case "GU":
		*e = CountryCodeAlpha2_Gu
	case "GT":
		*e = CountryCodeAlpha2_Gt
	case "GG":
		*e = CountryCodeAlpha2_Gg
	case "GN":
		*e = CountryCodeAlpha2_Gn
	case "GW":
		*e = CountryCodeAlpha2_Gw
	case "GY":
		*e = CountryCodeAlpha2_Gy
	case "HT":
		*e = CountryCodeAlpha2_Ht
	case "HM":
		*e = CountryCodeAlpha2_Hm
	case "VA":
		*e = CountryCodeAlpha2_Va
	case "HN":
		*e = CountryCodeAlpha2_Hn
	case "HK":
		*e = CountryCodeAlpha2_Hk
	case "HU":
		*e = CountryCodeAlpha2_Hu
	case "IS":
		*e = CountryCodeAlpha2_Is
	case "IN":
		*e = CountryCodeAlpha2_In
	case "ID":
		*e = CountryCodeAlpha2_ID
	case "IR":
		*e = CountryCodeAlpha2_Ir
	case "IQ":
		*e = CountryCodeAlpha2_Iq
	case "IE":
		*e = CountryCodeAlpha2_Ie
	case "IM":
		*e = CountryCodeAlpha2_Im
	case "IL":
		*e = CountryCodeAlpha2_Il
	case "IT":
		*e = CountryCodeAlpha2_It
	case "JM":
		*e = CountryCodeAlpha2_Jm
	case "JP":
		*e = CountryCodeAlpha2_Jp
	case "JE":
		*e = CountryCodeAlpha2_Je
	case "JO":
		*e = CountryCodeAlpha2_Jo
	case "KZ":
		*e = CountryCodeAlpha2_Kz
	case "KE":
		*e = CountryCodeAlpha2_Ke
	case "KI":
		*e = CountryCodeAlpha2_Ki
	case "KR":
		*e = CountryCodeAlpha2_Kr
	case "KP":
		*e = CountryCodeAlpha2_Kp
	case "KW":
		*e = CountryCodeAlpha2_Kw
	case "KG":
		*e = CountryCodeAlpha2_Kg
	case "LA":
		*e = CountryCodeAlpha2_La
	case "LV":
		*e = CountryCodeAlpha2_Lv
	case "LB":
		*e = CountryCodeAlpha2_Lb
	case "LS":
		*e = CountryCodeAlpha2_Ls
	case "LR":
		*e = CountryCodeAlpha2_Lr
	case "LY":
		*e = CountryCodeAlpha2_Ly
	case "LI":
		*e = CountryCodeAlpha2_Li
	case "LT":
		*e = CountryCodeAlpha2_Lt
	case "LU":
		*e = CountryCodeAlpha2_Lu
	case "MO":
		*e = CountryCodeAlpha2_Mo
	case "MK":
		*e = CountryCodeAlpha2_Mk
	case "MG":
		*e = CountryCodeAlpha2_Mg
	case "MW":
		*e = CountryCodeAlpha2_Mw
	case "MY":
		*e = CountryCodeAlpha2_My
	case "MV":
		*e = CountryCodeAlpha2_Mv
	case "ML":
		*e = CountryCodeAlpha2_Ml
	case "MT":
		*e = CountryCodeAlpha2_Mt
	case "MH":
		*e = CountryCodeAlpha2_Mh
	case "MQ":
		*e = CountryCodeAlpha2_Mq
	case "MR":
		*e = CountryCodeAlpha2_Mr
	case "MU":
		*e = CountryCodeAlpha2_Mu
	case "YT":
		*e = CountryCodeAlpha2_Yt
	case "MX":
		*e = CountryCodeAlpha2_Mx
	case "FM":
		*e = CountryCodeAlpha2_Fm
	case "MD":
		*e = CountryCodeAlpha2_Md
	case "MC":
		*e = CountryCodeAlpha2_Mc
	case "MN":
		*e = CountryCodeAlpha2_Mn
	case "ME":
		*e = CountryCodeAlpha2_Me
	case "MS":
		*e = CountryCodeAlpha2_Ms
	case "MA":
		*e = CountryCodeAlpha2_Ma
	case "MZ":
		*e = CountryCodeAlpha2_Mz
	case "MM":
		*e = CountryCodeAlpha2_Mm
	case "NA":
		*e = CountryCodeAlpha2_Na
	case "NR":
		*e = CountryCodeAlpha2_Nr
	case "NP":
		*e = CountryCodeAlpha2_Np
	case "NL":
		*e = CountryCodeAlpha2_Nl
	case "AN":
		*e = CountryCodeAlpha2_An
	case "NC":
		*e = CountryCodeAlpha2_Nc
	case "NZ":
		*e = CountryCodeAlpha2_Nz
	case "NI":
		*e = CountryCodeAlpha2_Ni
	case "NE":
		*e = CountryCodeAlpha2_Ne
	case "NG":
		*e = CountryCodeAlpha2_Ng
	case "NU":
		*e = CountryCodeAlpha2_Nu
	case "NF":
		*e = CountryCodeAlpha2_Nf
	case "MP":
		*e = CountryCodeAlpha2_Mp
	case "NO":
		*e = CountryCodeAlpha2_No
	case "OM":
		*e = CountryCodeAlpha2_Om
	case "PK":
		*e = CountryCodeAlpha2_Pk
	case "PW":
		*e = CountryCodeAlpha2_Pw
	case "PS":
		*e = CountryCodeAlpha2_Ps
	case "PA":
		*e = CountryCodeAlpha2_Pa
	case "PG":
		*e = CountryCodeAlpha2_Pg
	case "PY":
		*e = CountryCodeAlpha2_Py
	case "PE":
		*e = CountryCodeAlpha2_Pe
	case "PH":
		*e = CountryCodeAlpha2_Ph
	case "PN":
		*e = CountryCodeAlpha2_Pn
	case "PL":
		*e = CountryCodeAlpha2_Pl
	case "PT":
		*e = CountryCodeAlpha2_Pt
	case "PR":
		*e = CountryCodeAlpha2_Pr
	case "QA":
		*e = CountryCodeAlpha2_Qa
	case "RE":
		*e = CountryCodeAlpha2_Re
	case "RO":
		*e = CountryCodeAlpha2_Ro
	case "RU":
		*e = CountryCodeAlpha2_Ru
	case "RW":
		*e = CountryCodeAlpha2_Rw
	case "BL":
		*e = CountryCodeAlpha2_Bl
	case "SH":
		*e = CountryCodeAlpha2_Sh
	case "KN":
		*e = CountryCodeAlpha2_Kn
	case "LC":
		*e = CountryCodeAlpha2_Lc
	case "MF":
		*e = CountryCodeAlpha2_Mf
	case "PM":
		*e = CountryCodeAlpha2_Pm
	case "VC":
		*e = CountryCodeAlpha2_Vc
	case "WS":
		*e = CountryCodeAlpha2_Ws
	case "SM":
		*e = CountryCodeAlpha2_Sm
	case "ST":
		*e = CountryCodeAlpha2_St
	case "SA":
		*e = CountryCodeAlpha2_Sa
	case "SN":
		*e = CountryCodeAlpha2_Sn
	case "RS":
		*e = CountryCodeAlpha2_Rs
	case "SC":
		*e = CountryCodeAlpha2_Sc
	case "SL":
		*e = CountryCodeAlpha2_Sl
	case "SG":
		*e = CountryCodeAlpha2_Sg
	case "SK":
		*e = CountryCodeAlpha2_Sk
	case "SI":
		*e = CountryCodeAlpha2_Si
	case "SB":
		*e = CountryCodeAlpha2_Sb
	case "SO":
		*e = CountryCodeAlpha2_So
	case "ZA":
		*e = CountryCodeAlpha2_Za
	case "GS":
		*e = CountryCodeAlpha2_Gs
	case "ES":
		*e = CountryCodeAlpha2_Es
	case "LK":
		*e = CountryCodeAlpha2_Lk
	case "SD":
		*e = CountryCodeAlpha2_Sd
	case "SR":
		*e = CountryCodeAlpha2_Sr
	case "SJ":
		*e = CountryCodeAlpha2_Sj
	case "SZ":
		*e = CountryCodeAlpha2_Sz
	case "SE":
		*e = CountryCodeAlpha2_Se
	case "CH":
		*e = CountryCodeAlpha2_Ch
	case "SY":
		*e = CountryCodeAlpha2_Sy
	case "TW":
		*e = CountryCodeAlpha2_Tw
	case "TJ":
		*e = CountryCodeAlpha2_Tj
	case "TZ":
		*e = CountryCodeAlpha2_Tz
	case "TH":
		*e = CountryCodeAlpha2_Th
	case "TL":
		*e = CountryCodeAlpha2_Tl
	case "TG":
		*e = CountryCodeAlpha2_Tg
	case "TK":
		*e = CountryCodeAlpha2_Tk
	case "TO":
		*e = CountryCodeAlpha2_To
	case "TT":
		*e = CountryCodeAlpha2_Tt
	case "TN":
		*e = CountryCodeAlpha2_Tn
	case "TR":
		*e = CountryCodeAlpha2_Tr
	case "TM":
		*e = CountryCodeAlpha2_Tm
	case "TC":
		*e = CountryCodeAlpha2_Tc
	case "TV":
		*e = CountryCodeAlpha2_Tv
	case "UG":
		*e = CountryCodeAlpha2_Ug
	case "UA":
		*e = CountryCodeAlpha2_Ua
	case "AE":
		*e = CountryCodeAlpha2_Ae
	case "GB":
		*e = CountryCodeAlpha2_Gb
	case "US":
		*e = CountryCodeAlpha2_Us
	case "UM":
		*e = CountryCodeAlpha2_Um
	case "UY":
		*e = CountryCodeAlpha2_Uy
	case "UZ":
		*e = CountryCodeAlpha2_Uz
	case "VU":
		*e = CountryCodeAlpha2_Vu
	case "VE":
		*e = CountryCodeAlpha2_Ve
	case "VN":
		*e = CountryCodeAlpha2_Vn
	case "VG":
		*e = CountryCodeAlpha2_Vg
	case "VI":
		*e = CountryCodeAlpha2_Vi
	case "WF":
		*e = CountryCodeAlpha2_Wf
	case "EH":
		*e = CountryCodeAlpha2_Eh
	case "YE":
		*e = CountryCodeAlpha2_Ye
	case "ZM":
		*e = CountryCodeAlpha2_Zm
	case "ZW":
		*e = CountryCodeAlpha2_Zw
	case "SS":
		*e = CountryCodeAlpha2_Ss
	case "XK":
		*e = CountryCodeAlpha2_Xk
	case "BQ":
		*e = CountryCodeAlpha2_Bq
	default:
		return errors.New("jet: Invalid scan value '" + enumValue + "' for CountryCodeAlpha2 enum")
	}

	return nil
}

func (e CountryCodeAlpha2) String() string {
	return string(e)
}
