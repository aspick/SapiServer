var storage_path = WScript.Arguments(0);
var voice_index  = WScript.Arguments(1);
if (storage_path == undefined || voice_index == undefined) {
    WScript.Quit(1);
}

WScript.Echo(storage_path);

var tts    = WScript.CreateObject("SAPI.SpVoice");
var stream = WScript.CreateObject("SAPI.SpFileStream");

stream.open(storage_path + ".wav", 3);

tts.AudioOutputStream = stream;
tts.Voice = tts.GetVoices().Item(voice_index);

var sr      = new ActiveXObject("ADODB.Stream");
sr.Type     = 2     // text mode
sr.charset  = "utf-8";
sr.Open();
sr.LoadFromFile(storage_path + ".txt");

var message = sr.ReadText(-1);  // all line
sr.Close();
fs = null;

if (message.indexOf("</") == -1){
    tts.Speak(message);
}else{
    tts.Speak(message, 8);
}

WScript.StdOut.Write(message);
stream.close();
WScript.Quit(0);
