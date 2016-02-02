var tts    = WScript.CreateObject("SAPI.SpVoice");
voices = tts.GetVoices();

WScript.StdOut.Write("[");
for(var i = 0; i < voices.Count; i++){
    if(i > 0){
        WScript.StdOut.Write(",");
    }
    voice = voices.Item(i);
    WScript.StdOut.Write("{\"index\":" + i + ", ");
    WScript.StdOut.Write("\"name\":\"" + voice.GetAttribute("name") + "\", ");
    WScript.StdOut.Write("\"gender\":\"" + voice.GetAttribute("gender") + "\", ");
    WScript.StdOut.Write("\"language\":\"" + voice.GetAttribute("language") + "\", ");
    WScript.StdOut.Write("\"vendor\":\"" + voice.GetAttribute("vendor") + "\", ");
    WScript.StdOut.Write("\"age\":\"" + voice.GetAttribute("age") + "\", ");
    WScript.StdOut.Write("\"description\":\"" + voice.GetDescription() + "\"} ");

}
WScript.StdOut.Write("]");
