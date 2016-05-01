# Loverage

Loverage is a tool for golang projects that verbalizes a test coverage basing
on a testcases names.


### before

```
 λ  ~/go/src/github.com/kovetskiy/lorg: grep 'func Test' *_test.go -h | cut -d'(' -f1 | sed 's/func Test//'
```

```
NewFormat_ReturnsFormatWithDefaultFields
Format_ImplementsFormatterInterface
Format_GetPlaceholders_ReturnsPlaceholdersField
Format_SetPlaceholders_ChangesPlaceholdersField
Format_SetPlaceholder_ChangesPlaceholdersField
Format_Render_UsesPlaceholdersFieldForMatchingPlaceholders
Format_Render_PlaceholderRegexpMatching
Format_Reset_DesolatesReplacementAndUnsetsCompiledFlag
Format_Render_CallsSettedPlaceholders
Format_Render_CallsSettedPlaceholdersAndPassesLogLevel
Log_Fatal_ExitWithCode1
NewLog_ReturnsLogWithDefaultFields
Log_ImplementsLoggerInterface
Log_SetFormat_ChangesFormatField
Log_SetLevel_ChangesLevelField
Log_SetOutput_ChangesOutputField
Log_LoggingFunctions_CallsFormatRender
Log_LoggingFunctions_LogsRecordsWithSameLevelOrAbove
PlaceholderLevel_ReturnsLevelStringRepresentation
PlaceholderLine_ReturnsCallerLine
PlaceholderFile_ReturnsCallerFilenameInShortModeByDefault
PlaceholderFile_ReturnsCallerFilenameInShortModeIfValueIsShort
PlaceholderFile_ReturnsCallerFullFilenameInLongMode
PlaceholderTime_ReturnsTimestampIfValueIsTimestamp
PlaceholderTime_ReturnsTimeUsingDefaultLayoutIfValueNotSpecified
PlaceholderTime_ReturnsTimeUsingSpecifiedLayout
```

### using loverage

```
 λ  ~/go/src/github.com/kovetskiy/lorg: loverage
```

```
NewFormat                Returns Format With Default Fields
Format                   Implements Formatter Interface
Format.GetPlaceholders   Returns Placeholders Field
Format.SetPlaceholders   Changes Placeholders Field
Format.SetPlaceholder    Changes Placeholders Field
Format.Render            Uses Placeholders Field For Matching Placeholders
Format.Render            Placeholder Regexp Matching
Format.Reset             Desolates Replacement And Unsets Compiled Flag
Format.Render            Calls Setted Placeholders
Format.Render            Calls Setted Placeholders And Passes Log Level
Log.Fatal                Exit With Code1
NewLog                   Returns Log With Default Fields
Log                      Implements Logger Interface
Log.SetFormat            Changes Format Field
Log.SetLevel             Changes Level Field
Log.SetOutput            Changes Output Field
Log.LoggingFunctions     Calls Format Render
Log.LoggingFunctions     Logs Records With Same Level Or Above
PlaceholderLevel         Returns Level String Representation
PlaceholderLine          Returns Caller Line
PlaceholderFile          Returns Caller Filename In Short Mode By Default
PlaceholderFile          Returns Caller Filename In Short Mode If Value Is Short
PlaceholderFile          Returns Caller Full Filename In Long Mode
PlaceholderTime          Returns Timestamp If Value Is Timestamp
PlaceholderTime          Returns Time Using Default Layout If Value Not Specified
PlaceholderTime          Returns Time Using Specified Layout
```
