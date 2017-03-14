
# CreateMediaParams

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**name** | **String** | Name of media |  [optional]
**origin** | **String** | &#39;tts&#39;, &#39;file&#39; |  [optional]
**type** | **String** | &#39;hold_music&#39;, &#39;greeting&#39; |  [optional]
**ttsVoice** | **String** | &#39;allison&#39;, &#39;amy&#39;, &#39;belle&#39;, &#39;callie&#39;, &#39;callieq&#39;, &#39;dallas&#39;, &#39;damien&#39;, &#39;david&#39;, &#39;designerdave&#39;, &#39;diane&#39;, &#39;diesel&#39;, &#39;dog&#39;, &#39;duchess&#39;, &#39;duncan&#39;, &#39;emily&#39;, &#39;evilgenius&#39;, &#39;frank&#39;, &#39;french-fry&#39;, &#39;gregory&#39;, &#39;isabelle&#39;, &#39;jean-pierre&#39;, &#39;jerkface&#39;, &#39;katrin&#39;, &#39;kayla&#39;, &#39;kidaroo&#39;, &#39;lawrence&#39;, &#39;layo&#39;, &#39;linda&#39;, &#39;marta&#39;, &#39;matthias&#39;, &#39;miguel&#39;, &#39;millie&#39;, &#39;princess&#39;, &#39;ransomnote&#39;, &#39;robin&#39;, &#39;shouty&#39;, &#39;shygirl&#39;, &#39;tamika&#39;, &#39;tophat&#39;, &#39;vittoria&#39;, &#39;vixen&#39;, &#39;vlad&#39;, &#39;walter&#39;, &#39;whispery&#39;, &#39;william&#39;, &#39;wiseguy&#39;, &#39;zach&#39; |  [optional]
**ttsText** | **String** | Text used for text-to-speech conversion, maximum 800 characters |  [optional]
**isTemparary** | **String** | &#39;Y&#39;, &#39;N&#39;. Media file is temporary, will be deleted after a specified period. |  [optional]
**expirationDate** | **Integer** | If is_temporary is &#39;Y&#39;, media will be deleted after the specified time in seconds |  [optional]
**duration** | **Integer** | Length of media in seconds |  [optional]
**notes** | **String** | Notes about the media object |  [optional]
**randomized** | **String** | &#39;Y&#39;, &#39;N&#39;. Start playing the media file in random location, instead of from the beginning |  [optional]



