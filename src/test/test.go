package main

import (
	"crypto/sha256"
	"fmt"
	"encoding/hex"
	"github.com/Nik-U/pbc"
)


func main() {
	//str := "0x198750c02cdc32b35a991a5cbf240e8c97deb00e4a1a29cd3101766e67594eb6078b531b8b54289812d7f058efd9fa2a32cf611e53011d11e7ab03f9170b76d72002e9934c8e00a5abacb0dd2f04be31209c6b334a8f6eb9744c82056d4fe0f27b168babf1489f01b9bedfe80a9dc08af8723d8eedf4ae3a464e21732acca435f7200d83782dfe4d67810c100529fd18ac7c62c5c5b17aa97f1ecdddaba9bec3a9f0157af4494b2a350f546d9af3655b77efe0e7c86ba7c395f98c1670c93aae40922008f720cd0c68572c57833aeda0a647038ff2bf5c90e81b30391e8f2d226113470550c2ef0d82ac046818b6c7679e9194178e136dfe5f790e08b271961824c82620067ab876de4ba1ef1e1808b931761918c8d9c463f7ff9d6951b16d4b519594dc07eb0fb49a1283baaff6fb32e173e68635167f7b1782b0d9f5c40d360328b4812010b1628e7daa8bda5c006615803d3b76ccf4bc84cc742657bc2d4fce761260110d91685acd45a6f5ac514cb6f35a7ef9cc808bfd8804cef3773df2645f5916302016a8f08a9333e9f8462cb02c9ea766e7933c63cd9587166f6d07340365758be8102875cdd763c9b9d27db088fc6522bc8f27d32515e9deb3d34fad8666c8a00c201949e4e77478c2b5731b817df378c1369fc8d55c1986fc7083469d01fdd9df201783a1f06c560ec1185cdf50a77c04c6c821bc23af485d8510ffc41a7777f3380b38fbc34a92ab555944c5226e12834908161a3409cee6508cfa5dacb5068df60322b611a65814bac32511555046ef6e268074f06b589a06516c043aad36782d056bceabfa12cc57036666252f6d675b7cfe81f53e5f84845e026a44565c6cd11ce0ed4f2fbd2ee4ede9643c3dd3af1662043f1ee486b48f231a6f3b2a97be6a17a32cedad6779678450ff3c1643be80b5304256f10c8a470db7521977953af41f12b7088d0acf83d364ace5ba87348ad124e03f498606e13e17557e42444fef0fa119feb508f44d32cda90871005037fe60468570cc811f6bdd0523725ec34615e211fc9acb141a06037883ecf162c905182b2e36296b90a4834aeb493dab2d11d3cd95bbef47d3d06b4764499dba170a9378e51691f86f2b6aac43b5c7af0804c7fc4958333cbec9bd8c208e2f588a799f8ac56408966bb758eb418287a27d072c87802629af3718333876e3561981aa7832e7078363fd5f0000b885990c8100fb04121b33f645e926a1cddf7e33899da50643e69bf997cab44754de3e8d00351490dc7efba76b4e2092461e7c629da2dc31b74bc6860c3a1fb5b5035370dd1701637d6c54f2ea47bf748beec9d6cea3a27d084b70c85ad1a22d996cad2677b6e8d919e8f4d6ccc6f37cba0e7d629f10284dfafa891bae6506a045f316d2241e97515531687c769287cdacd6e5351c5fe4b304028ebccb0947bf3ce28f4f7008dfc69a2977ff22000756700b0252e33a734edc5c4fe85c5a0055332a08d658ccf4d2f3a85544cab1c1c0ef1cca7cf86b7f32b0d06ac0a9dfde27217b8ccba78d21dd77fec8fc75b9e1cb7e"
	str:="198750c02cdc32b35a991a5cbf240e8c97deb00e4a1a29cd3101766e67594eb6078b531b8b54289812d7f058efd9fa2a32cf611e53011d11e7ab03f9170b76d702e9934c8e00a5abacb0dd2f04be31209c6b334a8f6eb9744c82056d4fe0f27b168babf1489f01b9bedfe80a9dc08af8723d8eedf4ae3a464e21732acca435f70d83782dfe4d67810c100529fd18ac7c62c5c5b17aa97f1ecdddaba9bec3a9f0157af4494b2a350f546d9af3655b77efe0e7c86ba7c395f98c1670c93aae409208f720cd0c68572c57833aeda0a647038ff2bf5c90e81b30391e8f2d226113470550c2ef0d82ac046818b6c7679e9194178e136dfe5f790e08b271961824c826067ab876de4ba1ef1e1808b931761918c8d9c463f7ff9d6951b16d4b519594dc07eb0fb49a1283baaff6fb32e173e68635167f7b1782b0d9f5c40d360328b48110b1628e7daa8bda5c006615803d3b76ccf4bc84cc742657bc2d4fce761260110d91685acd45a6f5ac514cb6f35a7ef9cc808bfd8804cef3773df2645f59163016a8f08a9333e9f8462cb02c9ea766e7933c63cd9587166f6d07340365758be8102875cdd763c9b9d27db088fc6522bc8f27d32515e9deb3d34fad8666c8a00c1949e4e77478c2b5731b817df378c1369fc8d55c1986fc7083469d01fdd9df201783a1f06c560ec1185cdf50a77c04c6c821bc23af485d8510ffc41a7777f3380b38fbc34a92ab555944c5226e12834908161a3409cee6508cfa5dacb5068df60322b611a65814bac32511555046ef6e268074f06b589a06516c043aad36782d056bceabfa12cc57036666252f6d675b7cfe81f53e5f84845e026a44565c6cd11ce0ed4f2fbd2ee4ede9643c3dd3af1662043f1ee486b48f231a6f3b2a97be6a17a32cedad6779678450ff3c1643be80b5304256f10c8a470db7521977953af41f12b7088d0acf83d364ace5ba87348ad124e03f498606e13e17557e42444fef0fa119feb508f44d32cda90871005037fe60468570cc811f6bdd0523725ec34615e211fc9acb141a06037883ecf162c905182b2e36296b90a4834aeb493dab2d11d3cd95bbef47d3d06b4764499dba170a9378e51691f86f2b6aac43b5c7af0804c7fc4958333cbec9bd8c208e2f588a799f8ac56408966bb758eb418287a27d072c87802629af3718333876e3561981aa7832e7078363fd5f0000b885990c8100fb04121b33f645e926a1cddf7e33899da50643e69bf997cab44754de3e8d00351490dc7efba76b4e2092461e7c629da2dc31b74bc6860c3a1fb5b5035370dd1701637d6c54f2ea47bf748beec9d6cea3a27d084b70c85ad1a22d996cad2677b6e8d919e8f4d6ccc6f37cba0e7d629f10284dfafa891bae6506a045f316d2241e97515531687c769287cdacd6e5351c5fe4b304028ebccb0947bf3ce28f4f7008dfc69a2977ff22000756700b0252e33a734edc5c4fe85c5a0055332a08d658ccf4d2f3a85544cab1c1c0ef1cca7cf86b7f32b0d06ac0a9dfde27217b8ccba78d21dd77fec8fc75b9e1cb7e"
	input, _ := hex.DecodeString(str)
	//var input []byte = []byte(str)
	fmt.Println(len(input))
	fmt.Println(input[:10])
	//params:=pbc.GenerateA(160,256)
	params,_:=pbc.NewParamsFromString(`type a
q 84312369776863396118688570993543420635649923821062436188301550557781782044467
h 115377728800688073087534096588
r 730750818665452757176057050065048642452048576511
exp2 159
exp1 110
sign1 1
sign0 -1`)
	pairing:=params.NewPairing()	
	fmt.Println(params.String())

	g1:= pairing.NewG1().SetBytes(input[:64])
	//h:= pairing.NewG1().SetBytes(input[64:128])
	u:= pairing.NewG1().SetBytes(input[128:192])
	v:= pairing.NewG1().SetBytes(input[192:256])
	g2:= pairing.NewG2().SetBytes(input[256:320])
	w:= pairing.NewG2().SetBytes(input[320:384])
	ehw:= pairing.NewGT().SetBytes(input[384:448])
	ehg2:= pairing.NewGT().SetBytes(input[448:512])

	c1:=pairing.NewG1().SetBytes(input[512:576])
	c2:=pairing.NewG1().SetBytes(input[576:640])
	c3:=pairing.NewG1().SetBytes(input[640:704])
	t1:=pairing.NewG1().SetBytes(input[704:768])
	t2:=pairing.NewG1().SetBytes(input[768:832])
	t3:=pairing.NewG1().SetBytes(input[832:896])
	c:=pairing.NewZr().SetBytes(input[896:916])
	salpha:=pairing.NewZr().SetBytes(input[916:936])
	sbeta:=pairing.NewZr().SetBytes(input[936:956])
	sa:=pairing.NewZr().SetBytes(input[956:976])
	sx:=pairing.NewZr().SetBytes(input[976:996])
	sdelta1:=pairing.NewZr().SetBytes(input[996:1016])
	sdelta2:=pairing.NewZr().SetBytes(input[1016:1036])
	h_:=pairing.NewG1().SetBytes(input[1036:])
	r1 :=pairing.NewG1().Mul(pairing.NewG1().PowZn(u,salpha),pairing.NewG1().PowZn(t1,pairing.NewZr().Neg(c)))
	r2 :=pairing.NewG1().Mul(pairing.NewG1().PowZn(v,sbeta),pairing.NewG1().PowZn(t2,pairing.NewZr().Neg(c)))
	//******************************************
	temp1:= pairing.NewGT().Pair(t3, g2)
	r3_e1:=pairing.NewGT().PowZn(temp1,sa)
	uuu:=pairing.NewZr().Neg(salpha)
	www:=pairing.NewZr().Neg(sbeta)
	xxx:=pairing.NewZr().Add(uuu,www)
	r3_e2:=pairing.NewGT().PowZn(ehw,xxx)
	uuu1:=pairing.NewZr().Neg(sdelta1)
	www1:=pairing.NewZr().Neg(sdelta2)
	xxx1:=pairing.NewZr().Add(uuu1,www1)
	r3_e3:=pairing.NewGT().PowZn(ehg2,xxx1)
	eh3g2:=pairing.NewGT().Pair(h_,g2)
	r3_e4:=pairing.NewGT().PowZn(eh3g2,sx)
	r3_tep:=pairing.NewGT().Mul(pairing.NewGT().Mul(r3_e1,r3_e2),pairing.NewGT().Mul(r3_e3,r3_e4))
	yyy:=pairing.NewGT().Pair(t3,w)
	ggg:=pairing.NewGT().Pair(g1,g2)
	hhh:=pairing.NewGT().Invert(ggg)
	r3:=pairing.NewGT().Mul(r3_tep,pairing.NewGT().PowZn(pairing.NewGT().Mul(yyy,hhh),c))
	//*******************************************
	tt_temp2:=pairing.NewG1().PowZn(t1,sa)
	tt_temp:=pairing.NewZr().Neg(sdelta1)
	tt:=pairing.NewG1().PowZn(u,tt_temp)
	r4:=pairing.NewG1().Mul(tt,tt_temp2)
	rr_temp2:=pairing.NewG1().PowZn(t2,sa)
	rr_temp:=pairing.NewZr().Neg(sdelta2)
	rr:=pairing.NewG1().PowZn(v,rr_temp)
	r5:=pairing.NewG1().Mul(rr,rr_temp2)
	var  s  string
	s+=t1.String()
	s+=t2.String()
	s+=t3.String()
	s+=r1.String()
	s+=r2.String()
	s+=r3.String()
	s+=r4.String()
	s+=r5.String()
	s+=c1.String()
	s+=c2.String()
	s+=c3.String()
	c_:= pairing.NewZr().SetFromStringHash(s,sha256.New())
	fmt.Printf("g1:%02x \n",g1.Bytes())
   //fmt.Printf("h:%02x \n",h.Bytes())
   fmt.Printf("u:%02x \n",u.Bytes())
   fmt.Printf("v:%02x \n",v.Bytes())
   fmt.Printf("g2:%02x \n",g2.Bytes())
   fmt.Printf("w:%02x \n",w.Bytes())
   fmt.Printf("ehw:%02x \n",ehw.Bytes())
   fmt.Printf("ehg2:%02x \n",ehg2.Bytes())
   //fmt.Printf("minusEg1g2:%02x \n",priv.minusEg1g2.Bytes())
   fmt.Printf("sig.c1:%02x \n",c1.Bytes())
        fmt.Printf("sig.c2:%02x \n",c2.Bytes())
        fmt.Printf("sig.c3:%02x \n",c3.Bytes())
        fmt.Printf("sig.c:%02x \n",c.Bytes())
        fmt.Printf("sig.t1:%02x \n",t1.Bytes())
        fmt.Printf("sig.t2:%02x \n",t2.Bytes())
        fmt.Printf("sig.t3:%02x \n",t3.Bytes())
        fmt.Printf("sig.salpha:%02x \n",salpha.Bytes())
        fmt.Printf("sig.sbeta:%02x \n",sbeta.Bytes())
        fmt.Printf("sig.sa:%02x \n",sa.Bytes())
        fmt.Printf("sig.sx:%02x \n",sx.Bytes())
        fmt.Printf("sig.sdelta1:%02x \n",sdelta1.Bytes())
        fmt.Printf("sig.sdelta2:%02x \n",sdelta2.Bytes())
	fmt.Printf("c_:%02x \n",c_.Bytes())

	fmt.Printf("r1:%02x \n",r1.Bytes())
	fmt.Printf("r2:%02x \n",r2.Bytes())
	fmt.Printf("r3:%02x \n",r3.Bytes())
	fmt.Printf("r4:%02x \n",r4.Bytes())
	fmt.Printf("r5:%02x \n",r5.Bytes())

	if c_.Equals(c){
		//true32Byte := []byte{1}
		fmt.Println(0)
		//return   true32Byte
	}else{
		//false32Byte := []byte{0}
		fmt.Println(1)
		//return  false32Byte
	}
}
