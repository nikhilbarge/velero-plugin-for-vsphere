/*
Copyright the Velero contributors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by crds_generate.go; DO NOT EDIT.

package crds

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"

	apiextinstall "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/install"
	apiextv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
)

var rawCRDs = [][]byte{
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcWώ\xdb6\x13\xbf\xfb)\x06\xf9\x0e\xfb\x15\x88\xb4\bZ\x14\x85o\xcdn\v,\x9a\x04\x8b8\xd8K\xd1\x03%\x8emv)\x0e\xcb!\xbdu\x9f\xbe\x18R\xb2eY\xeb\xf8\x10T7\r\xe7\xff\xfcf8\\TU\xb5P\xde<a`Cn\t\xca\x1b\xfc;\xa2\x93?\xae\x9f\u007f\xe2\xda\xd0\xed\xee]\x83Q\xbd[<\x1b\xa7\x97p\x978R\xf7\x19\x99Rh\xf1\x1e\xd7ƙh\xc8-:\x8cJ\xab\xa8\x96\v\x00\xe5\x1cE%d\x96_\x80\x96\\\fd-\x86j\x83\xae~N\r6\xc9X\x8d![\x18\xec\xff?\xb9gG/\xee\xbb\x05@\x1b0k\xf8b:\xe4\xa8:\xbf\x04\x97\xac]\x008\xd5\xe1\x124\xbd8KJs\xbdC\x8b\x81\xbcM\x1b\xe3jC\v\xf6؊\xd9M\xa0\xe4\x970=.\nz\xc7JP\xf7\xbd\xaeL\xb2\x86\xe3o'\xe4\x0f\x86c>\xf26\x05eG\xb63\x95\x8d\xdb$\xab\u0091\xbe\x00\xe0\x96<.\xe1͛\x05\xc0NY\xa3s0\xc5(yt??><}\xbfj\xb7ةB\x04\xd0\xc8m0>\xf3\x1dl\xf7\xd4\x06A\xf5\x91T%\x14\bȑ\x02\xf6\xc2>\x90\xc7\x10\xcd\x10\x98|\xa3\xd2\x1eh\x1337\xe2G\xe1\x01-\xc5D\x86\xb8E\xe8K\x82\x1a8\xfb\b\xb4\x86\xb85\f\x01}@FW\xca;R\v¢\x1cP\xf3'\xb6\xb1\x86\x15\x06Q\x02\xbc\xa5d\xb5 `\x87!B\xc0\x966\xce\xfcs\xd0\xcc\x10)\x9b\xb4*b\x9f\xe7\xe13.bp\xcaJ\x06\x13\xbe\x05\xe54tj\x0f\x01\xc5\x06$7ҖY\xb8\x86\x8f\x14\x10\x8c[\xd3\x12\xb61z^\xde\xdenL\x1c\xc0\xdcR\xd7%g\xe2\xfe6C\xd24)R\xe0[\x8d;\xb4\xb7ʛ*\xfb\xe92t\xebN\xff/\xf4@监cq/\xa5\xe5\x18\x8c\xdb\x1c\xc8\x19I\xaf\xa6Y\x00\x05\x86A\xf5b\xc5\xddc6\x85$I\xf8\xfc\xcb\xea\v\fFs\xc6OS\x9c\x93{\x14\xe3c\x9e%/ƭ1\x94:\xad\x03uY#:\xedɸ\x98\u007fZkН\xe6\x98Sә(\x85\xfd+!G)G\rw\xb9\x83\xa1AH^\xab\x88\xba\x86\a\aw\xaaC{\xa7\x18\xbfu\x96%\xa1\\I\x06\xbf\x9e\xe7\xf1\x9c9e,\xc99\x90\x8710[\x90\x95\xc7Vꑓ\x92G\xda1\xeb\"8\x92\x9b\xeb,\xf9\x1a\xd5>'\xbf\x8a\x14\xd4\x06?P;j\xf0W\x8c\xbe\x9f\x93\x18\xbc\x90\xa1T\xba\f{\xd5\xc0\x85s\xa2\x12\xc0\x0e\xa2/[\f\x98%d\x18\x98V\x80Al\"\x85\xbd\xa8\xcd\xf3A\xd7\x13\xf9٬\xca\xe7H\xe3E\xff?\x91\xc69wE\x10\xe2V\x15\x8c=RFzH\xce\t\xaa\xc9]\xed\x00;\xe5yK\xf1\xe1\xfe\xa2\x1b\xab\x03\xdb\xe0\x8c\xd1\x02\xa6\xb5\xc1\x00k\n\x994\xe8\x1a|\xf4\xa4\xcf\xf2\xb8#\x9b:\xbcڽ\xc2>X\xbf\xaa\xe2O\xb3\"\xe2\xb6\x0f\xb43\x1auv81\x06)\xa6\x13W\xf7\xf0\xa2&=\x9a\xf5\x9a\xf5\x1a\x03\xbax(\u007f\x9e\x9c\xb9\xc6edg[\xc0ʕ\xc0\xf3\xb0Ĩ@\xba\xe5\xca e\b\x98\x80'\x83\xac\x9aG\xfa\t\x87 \xe0\xab\xfd\x18UL\xfcjG\x0e\x17\xde*\xb3\x1d{3䘋p\xbeb\x0e\x9c\xf5\x15M\xdaR\xe7-\x9e\xee\x11\x97\nvwΟ\xaf\xac\xa0\x8b?\xd1tr\x13\xf7\r\xfa\xa2x\xb0p\xdeg0RV\xe4\xf2\xf5)\xbaP\x03\xee\xd0\x019X+cQ\xf7\n\xb9\x9eʜ\xe9\x1c\xebhp-\xc5O^\xb21\xdc \xbdk\xc35\xfcEz!_\x117\xfc\xaa\xc6\xc4=\x10g\xc2\xe7\x89\xc0\x9aB\xa7\xe2RP\x85ՌBY\xd2Tcq\t1\xa4\xe9\u1afd\xd5!\xb3\xda\\\x1e?\x1f\vO\xb9F{\x01P\r\xa52w\x86\xe5\xeb\x86{\xb4\\\xdd\xd8~\xab\xf8\xb2\xe9G\xe1\x98\xc3\xe4a\x06\u0381R>t\xa9\x9b\xaa\xae\xe0\x13\xbe\x9c\xd1\x1e\xdcc\xa0M@\x9ef\xbc\x1aʂ\xd3\x11V\xc1\xaf\x19>W\a\xda\x1b\xb8\x1ck\xcf\x04[\xb2\x03\xea)*\v.u\r\x06\t\xb8\xd9G\xe4!\xf22w\xce`%\xe3g\x9c\xad\xa3t\xbf\xb8B\xf2E\x91\x80\xd40\xb4\xca\xe5\x85C\xb0\x18\t\xb4ao\xd5\xfeL\xef\x10C\xde>\x04\x8a\xd2*G\x14\f\xcae\x1c\xe4\xb3iA^\x9b\x14\xf2ew\xeeɝ\x81\x01F\xc07.\xfe\xf8\xc3\xccyɺ\xac\xac\x1b\fg\xe79\x85\xefE\xff\xb7\xd6=;l\xa1\f\xdc\x10\xaf\x1b{\xab\x13֯L\xbc\xacWV\xc2\x15z\x15T<\xaf}^>\xef\xa6\x0f\xb8\xb7\xf2Rj\xfbM\xa1\xec\xbf\xedV\xb9\x8d@\xe9\xf0\x9c\xc9h8\xd7x2\xc2NF֩\xeb\xffŴ\x9aI\xf8\xf4ެ\xc6K\xe4\x84\xff\xf0\xdaݽ;\xfeeTT\xfd\xd36\x1f\xf41\xeb\x91g\xfd>\xd8S\x8e\xf7\xa9j[\xf4\x11\xf5\xa7\xe9\xbb6\xbf>\x8f\xcf\xd6\xfcےӦ<\xcd\xe1\xf7?\x16\xd0o\x89O\x83\x1fB\xfc7\x00\x00\xff\xff0\x93\xee\xef\x19\x10\x00\x00"),
	[]byte("\x1f\x8b\b\x00\x00\x00\x00\x00\x00\xff\xbcW_\x8f\xdb6\f\u007fϧ \xba\x87ۀڇb\xc30\xe4m\xbdn\xc0amqh\xba{\x19\xf6 [L\xa2\x9dLi\xa2\x94[\xf6\xe9\aJv\xe28\xbe4\x0f\xc5\xfcf\x8a\xff\xf9#E-\xaa\xaaZ(o\x1e1\xb0q\xb4\x04\xe5\r\xfe\x13\x91\xe4\x8f맟\xb86\xeev\xf7\xa6\xc1\xa8\xde,\x9e\f\xe9%\xdc%\x8e\xae\xfb\x84\xecRh\xf1\x1d\xae\r\x99h\x1c-:\x8cJ\xab\xa8\x96\v\x00E\xe4\xa2\x122\xcb/@\xeb(\x06g-\x86j\x83T?\xa5\x06\x9bd\xacƐ-\f\xf6\xbfM\xf4D\ue67e[\x00\xb4\x01\xb3\x86ϦC\x8e\xaa\xf3K\xa0d\xed\x02\x80T\x87KH\xde:\xa5\xb9ޡ\xc5\xe0\xbcM\x1bC\xb5q\v\xf6؊\xd1Mp\xc9/az\\\xc4{\xb7JH\xbfgM\x99`\r\xc7\xdfF\xc4\xf7\x86c>\xf06\x05e\x0fV3\x8d\rm\x92Ua\xa0.\x00\xb8u\x1e\x97\xf0\xea\xd5\x02`\xa7\xac\xd19\x84b\xccy\xa4\x9f\x1f\xee\x1f\xbf_\xb5[\xecT!\x02h\xe46\x18\x9f\xf9z\xab=\xadAP\xbd\xffU\t\x00\x1a\xd5>%\xdfK\xfa\xe0<\x86h\x86h\xe4\x1bU\xf3@\x9bظ\x11'\n\x0fh\xa9\x1f2\xc4-B_\x05\xd4\xc0\xd9Apk\x88[\xc3\x10\xd0\ad\xa4RёZ\x10\x16E\xe0\x9a\xbf\xb0\x8d5\xac0\x88\x12\xe0\xadKVK\xd1w\x18\"\x04l݆̿\a\xcd\f\xd1e\x93VE\xec\x13<|\x86\"\x06RVҗ\xf05(\xd2Щ=\x04\x14\x1b\x90h\xa4-\xb3p\r\x1f\\@0\xb4vK\xd8\xc6\xe8yy{\xbb1q\xc0o\xeb\xba.\x91\x89\xfbیBӤ\xe8\x02\xdfjܡ\xbdU\xdeT\xd9O\xcah\xad;\xfdM\xe8\xb1\xcd7#\xc7\xe2^\xea\xca1\x18\xda\x1c\xc8\x19>/\xa6Yp\x04\x86A\xf5b\xc5\xddc6\x85$I\xf8\xf4\xcb\xea3\fFs\xc6OS\x9c\x93{\x14\xe3c\x9e%/\x86\xd6\x18J\x9d\xd6\xc1uY#\x92\xf6\xceP\xcc?\xad5H\xa79\xe6\xd4t&Ja\xffN\xc8Q\xcaQ\xc3]nZh\x10\x92\xd7*\xa2\xae\xe1\x9e\xe0Nuh\xef\x14\xe3\xd7β$\x94+\xc9\xe0\x97\xf3<\x1e-\xa7\x8c%9\a\xf2\xd0\xfb\xb3\x05Yyl\xa5\x1e9)y\x8a\x1d\xb3.\x82#\xb9\xb9Β\xaf\xb4\xdf*\xba\xa06\xf8\u07b5\xa3\xee~\xc1\xe8\xdb9\x89\xc1\v\x99D\xa5˰W\r\\8'*\x01\xec \xfa\xbcŀY\" G\xd3\n0\x1c\x9b\xe8\xc2^Ԋ<\xeaz\"?\x9bU\xf9\xc8i\xbc\xe8\xffG\xa7q\xce]\x11\x84\xb8U\x05c\x0f.#=$\"A\xb5\xa3\xab\x1d`R\x9e\xb7.\u07bf\xbb\xe8\xc6\xea\xc068c\xb4\x80im0\xc0څL\x1at\r>z\xa7\xcf\xf2\xb8s6ux\xb5{\x85}\xb0~U\xc5\x1fgE\xc4m\x1f\xdc\xceh\xd4\xd9\xe1\xc4\x18\xa4\x98$\xae\xee\xe1YMz4\xeb5\xeb5\x06\xa4x(\u007f\x9e\x9c\xb9\xc6edg[\xc0\x8aJ\xe0yXbT \xddre\x902\x04L\xc0\x93AV\xcd#\xfd\x84C\x10\xf0\xc5~\x8c*&~\xb1#\xcbm\xb7\xcaL\xc7\xce\f9\xe2\"\x9a/\x98\x9e\xaf\xbe\xa2A[\xd7y\x8b\xa7kåbݝ\xf3\xe7\xeb*\xe8\xe2M4\x9d\\\xc2}s>+\x1e,\x9c\xf7\x18\x8c\x94\x15\xb9|u\x8a.Ԁ;$p\x04ke,\xea^!\xd7S\x993\x9dc\x1d\r\xae\xa5\xf0e\xd9\x18n\x8f\u07b5\xe1\n\xfe,}\x90\xaf\x87\x1b~Qc\xe2\x1e\x843\xe1\xf3D`\xedB\xa7\xe2R\x10\x85ՌB\xd9\xc9Tcq\t1\xa4\xe9\xe1\x8b}\xd5!\xb3\xda\\\x1e=\x1f\nO\xb9B{\x01P\x8dKe\xe6\x944\xdcp\x8f\x94\xab[\xdao\x15_6\xfc \x1csx<L\xbfs@ʇ\x94\xba\xa9\xe2\n>\xe2\xf3\x19\xed\x9e\x1e\x82\xdb\x04\xe4i\xb6\xab\xa1$8\x1d]\x15\xfc\x9a\xa1su\x98\xbd\x81ˑ\xf6L\xb0uv@\xbc\x8b\xca\x02\xa5\xae\xc1 \xe16\xfb\x88<\xc4]\xe6\xcd\x19\xa4d\xec\x8cs5\x92V\xed\x13jH\xbe\xe8\x11|\x1a\x86VQ\xde3\x04\x86с6\xec\xadڟ\xa9\x1dB\xc8K\x87\xa0P\xba\xe4\b\x80\x01\xf9\x1eC>\x9a\x96\xe3\xa5\x19!_\xf6杣3 \xc0\b\xf2\x86\xe2\x8f?̜\x97\x9cˢ\xba\xc1pv\x9e\x13\xf8V\xf4\u007fmݳ#\x16ʘ\r\U0007a077:a\xfd¬\xcbze\x11\\\xa1WA\xc5\xf3\xca\xe7\x95\xf3n\xfaR{-O\xa3\xb6\xdf\x0f\xca\xd6\xdbn\x15m\x04H\x94\xf7\x16\x17\n\x18\xce5\x9e\f\xaf\x93au\xea\xfa\xff1\xa7f\x12>\xbd-\xab\xf1\xea8\xe1?<kwo\x8e\u007f\x19\x15U\xff\x86\xcd\a}\xccz\xe4Y\xbf\x05\xf6\x94\xe3-\xaa\xda\x16}D\xfdq\xfa\x84\xcd\x0f\xce\xe3+5\xff\xb6\x8e\xb4)op\xf8\xe3\xcf\x05\xf4\xbb\xe1\xe3\xe0\x87\x10\xff\v\x00\x00\xff\xff\xad\x90϶\x02\x10\x00\x00"),
}

var CRDs = crds()

func crds() []*apiextv1beta1.CustomResourceDefinition {
	apiextinstall.Install(scheme.Scheme)
	decode := scheme.Codecs.UniversalDeserializer().Decode
	var objs []*apiextv1beta1.CustomResourceDefinition
	for _, crd := range rawCRDs {
		gzr, err := gzip.NewReader(bytes.NewReader(crd))
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(gzr)
		if err != nil {
			panic(err)
		}
		gzr.Close()

		obj, _, err := decode(bytes, nil, nil)
		if err != nil {
			panic(err)
		}
		objs = append(objs, obj.(*apiextv1beta1.CustomResourceDefinition))
	}
	return objs
}
