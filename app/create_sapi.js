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

var fs      = new ActiveXObject("Scripting.FileSystemObject");
var file    = fs.OpenTextFile(storage_path + ".txt", 1, true, 0);
var message = file.ReadAll();
file.Close();
fs = null;

if (message.indexOf("</") == -1){
    tts.Speak(message);
}else{
    tts.Speak(message, 8);
}

WScript.StdOut.Write(message);
stream.close();
WScript.Quit(0);
