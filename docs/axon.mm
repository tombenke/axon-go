<map version="freeplane 1.3.0">
<!--To view this file, download free mind mapping software Freeplane from http://freeplane.sourceforge.net -->
<node TEXT="Axon" FOLDED="false" ID="ID_360047530" CREATED="1458311830319" MODIFIED="1605977541846"><hook NAME="MapStyle">
    <properties show_icon_for_attributes="true" show_note_icons="true"/>

<map_styles>
<stylenode LOCALIZED_TEXT="styles.root_node">
<stylenode LOCALIZED_TEXT="styles.predefined" POSITION="right">
<stylenode LOCALIZED_TEXT="default" MAX_WIDTH="600" COLOR="#000000" STYLE="as_parent">
<font NAME="SansSerif" SIZE="10" BOLD="false" ITALIC="false"/>
</stylenode>
<stylenode LOCALIZED_TEXT="defaultstyle.details"/>
<stylenode LOCALIZED_TEXT="defaultstyle.note"/>
<stylenode LOCALIZED_TEXT="defaultstyle.floating">
<edge STYLE="hide_edge"/>
<cloud COLOR="#f0f0f0" SHAPE="ROUND_RECT"/>
</stylenode>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.user-defined" POSITION="right">
<stylenode LOCALIZED_TEXT="styles.topic" COLOR="#18898b" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
<cloud COLOR="#f0f0f0" SHAPE="ARC"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subtopic" COLOR="#cc3300" STYLE="fork">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.subsubtopic" COLOR="#669900">
<font NAME="Liberation Sans" SIZE="10" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.important">
<icon BUILTIN="yes"/>
<cloud COLOR="#f0f0f0" SHAPE="ARC"/>
</stylenode>
<stylenode TEXT="lsyh" COLOR="#990000">
<font NAME="SansSerif" SIZE="10" BOLD="true"/>
<edge COLOR="#808080"/>
</stylenode>
<stylenode TEXT="home" COLOR="#215800">
<font SIZE="10" BOLD="true"/>
</stylenode>
<stylenode TEXT="ohome" COLOR="#120088" BACKGROUND_COLOR="#fdfd51">
<font NAME="SansSerif" SIZE="12" BOLD="true"/>
<cloud COLOR="#fdfd51" SHAPE="ROUND_RECT"/>
<edge COLOR="#808080"/>
</stylenode>
<stylenode TEXT="activity" COLOR="#102292" BACKGROUND_COLOR="#88d8d9" STYLE="bubble">
<font SIZE="12" BOLD="true"/>
</stylenode>
<stylenode TEXT="goal" BACKGROUND_COLOR="#b6e174">
<font BOLD="true"/>
</stylenode>
<stylenode TEXT="activityDetailHiPrio" BACKGROUND_COLOR="#e9f664">
<font BOLD="true"/>
</stylenode>
<stylenode TEXT="activityDetailLoPrio" BACKGROUND_COLOR="#f4f8d3"/>
</stylenode>
<stylenode LOCALIZED_TEXT="styles.AutomaticLayout" POSITION="right">
<stylenode LOCALIZED_TEXT="AutomaticLayout.level.root" COLOR="#000000">
<font SIZE="18"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,1" COLOR="#0033ff">
<font NAME="Monospaced" BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,2" COLOR="#00b439">
<font BOLD="true"/>
<edge COLOR="#808080"/>
<cloud COLOR="#f0f0f0" SHAPE="ARC"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,3" COLOR="#990000">
<font BOLD="true"/>
</stylenode>
<stylenode LOCALIZED_TEXT="AutomaticLayout.level,4" COLOR="#111111">
<font NAME="SansSerif" SIZE="10"/>
</stylenode>
</stylenode>
</stylenode>
</map_styles>
</hook>
<node TEXT="notes:" POSITION="right" ID="ID_1775331280" CREATED="1458311884012" MODIFIED="1458312182392">
<node TEXT="design" STYLE_REF="home" FOLDED="true" ID="ID_71273065" CREATED="1590331262082" MODIFIED="1605977742379">
<icon BUILTIN="idea"/>
<node TEXT="node" ID="ID_1454386825" CREATED="1590331525634" MODIFIED="1590331528087">
<node TEXT="azonos&#xed;t&#xe1;sa" ID="ID_748564640" CREATED="1590331280387" MODIFIED="1590331540496">
<node TEXT="nuuid" ID="ID_1757729395" CREATED="1590331289354" MODIFIED="1590331291748">
<node TEXT="A node-oknak kell egy unique UUID" ID="ID_89658497" CREATED="1590331268524" MODIFIED="1590331279813"/>
<node TEXT="minden" ID="ID_465440454" CREATED="1590331286786" MODIFIED="1590331288296"/>
</node>
<node TEXT="name" ID="ID_198975985" CREATED="1590331515389" MODIFIED="1590331516417">
<node TEXT="symbolic (functional) name" ID="ID_208172909" CREATED="1590331292083" MODIFIED="1590331512311"/>
<node TEXT="namespace" ID="ID_1383041151" CREATED="1590331299219" MODIFIED="1590331301191"/>
<node TEXT="remapping" ID="ID_468695972" CREATED="1590331296612" MODIFIED="1590331298967"/>
</node>
</node>
<node TEXT="tulajdons&#xe1;gok" ID="ID_1065001508" CREATED="1590331541315" MODIFIED="1590331544817">
<node TEXT="subscriptions" ID="ID_956445584" CREATED="1590331545298" MODIFIED="1590331647924">
<node TEXT="(inbound) channels" ID="ID_798202474" CREATED="1590331651011" MODIFIED="1590331662112"/>
<node TEXT="accepted messages" ID="ID_1244781878" CREATED="1590331565411" MODIFIED="1590331686759">
<node TEXT="message-type / content-type" ID="ID_374300147" CREATED="1590331593299" MODIFIED="1590331605109"/>
</node>
</node>
<node TEXT="publications" ID="ID_690731317" CREATED="1590331557473" MODIFIED="1590331671941">
<node TEXT="(outbound) channels" ID="ID_1787983959" CREATED="1590331672884" MODIFIED="1590331677358"/>
<node TEXT="delivered (???)  messages" ID="ID_1076975356" CREATED="1590331678243" MODIFIED="1590331699623">
<node TEXT="message-type / content-type" ID="ID_590993663" CREATED="1590331593299" MODIFIED="1590331605109"/>
</node>
</node>
<node TEXT="services" ID="ID_238828919" CREATED="1590331769096" MODIFIED="1590331771375">
<node TEXT="RPC-like service endpoints" ID="ID_850668439" CREATED="1590331786147" MODIFIED="1590331792594"/>
</node>
</node>
</node>
<node TEXT="messages" ID="ID_450854447" CREATED="1590331343315" MODIFIED="1590331345855">
<node TEXT="k&#xf6;zponti message schema &#xe9;s fixture store" ID="ID_1294168953" CREATED="1590331346240" MODIFIED="1590331377253">
<node TEXT="standard messages" ID="ID_1517317311" CREATED="1590331879459" MODIFIED="1590331883359"/>
<node TEXT="application-f&#xfc;gg&#x151; gy&#x171;jtem&#xe9;nyek" ID="ID_1008175819" CREATED="1590331884356" MODIFIED="1590331901141">
<node TEXT="GPIO" ID="ID_584209008" CREATED="1590331910625" MODIFIED="1590331918614"/>
<node TEXT="stepper" ID="ID_1139130904" CREATED="1590331919236" MODIFIED="1590331920795"/>
<node TEXT="DSIM" ID="ID_583402795" CREATED="1590331921433" MODIFIED="1590331959904"/>
</node>
</node>
<node TEXT="" ID="ID_1566806357" CREATED="1590331387747" MODIFIED="1590331387747">
<node TEXT="standard messages" ID="ID_1262047683" CREATED="1590331379300" MODIFIED="1590331383040"/>
<node TEXT="compound messages" ID="ID_796401443" CREATED="1590331383279" MODIFIED="1590331386866"/>
<node TEXT="custom messages" ID="ID_1957858160" CREATED="1590331391652" MODIFIED="1590331395361"/>
</node>
<node TEXT="generic structure:" FOLDED="true" ID="ID_1043555185" CREATED="1590331406787" MODIFIED="1590911627782">
<node TEXT="notes:" ID="ID_1779398158" CREATED="1590911702659" MODIFIED="1590911704461">
<node TEXT="Az &#xfc;zeneteknek minden esetben van egy { header, body } strukt&#xfa;r&#xe1;ja, a bels&#x151; &#xe1;br&#xe1;zol&#xe1;s szerint" ID="ID_1874245772" CREATED="1590911585910" MODIFIED="1590911701818"/>
<node TEXT="Az alap&#xe9;rtelmezett content type" ID="ID_1854754781" CREATED="1590911710157" MODIFIED="1590911723525"/>
</node>
<node TEXT="header" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1357928537" CREATED="1590331403971" MODIFIED="1590911707428">
<node TEXT="// meta info" ID="ID_1999031020" CREATED="1590331412979" MODIFIED="1590331418557"/>
<node TEXT="message-type" ID="ID_1421351549" CREATED="1590331423530" MODIFIED="1590331432868">
<node TEXT="mandatory" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1831500320" CREATED="1590911637212" MODIFIED="1590911681682"/>
<node TEXT="default:" ID="ID_1958623615" CREATED="1590911725592" MODIFIED="1590911727066">
<node TEXT="axon:any" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_149050148" CREATED="1590911727932" MODIFIED="1590911750878"/>
</node>
</node>
<node TEXT="content-type" ID="ID_374017291" CREATED="1590331427103" MODIFIED="1590331430325">
<node TEXT="mandatory" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1146998168" CREATED="1590911640331" MODIFIED="1590911748246"/>
<node TEXT="default:" ID="ID_1528374754" CREATED="1590911738964" MODIFIED="1590911741046">
<node TEXT="application/JSON" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_221205353" CREATED="1590911741616" MODIFIED="1590911747989"/>
</node>
</node>
<node TEXT="other (HTTP-like) headers" ID="ID_1597970213" CREATED="1590911651777" MODIFIED="1590911659694">
<node TEXT="optional" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1303539" CREATED="1590911660125" MODIFIED="1590911693059"/>
<node TEXT="egy&#xe9;b HTTP resp. header-eket megvizsg&#xe1;lni" ID="ID_1089768946" CREATED="1590331433593" MODIFIED="1590331450940">
<node TEXT="" ID="ID_1320812632" CREATED="1590334915094" MODIFIED="1590334915094">
<node TEXT="all" ID="ID_1744388056" CREATED="1590334916715" MODIFIED="1590334930874">
<node TEXT="Content-Length" ID="ID_375903132" CREATED="1590334931481" MODIFIED="1590334935854"/>
<node TEXT="Content-Language" ID="ID_971608017" CREATED="1590334936249" MODIFIED="1590334939030"/>
<node TEXT="Content-Encoding" ID="ID_1413087120" CREATED="1590334944154" MODIFIED="1590335166687">
<font BOLD="true"/>
</node>
<node TEXT="Content-Location" ID="ID_360174140" CREATED="1590334967137" MODIFIED="1590334973050"/>
<node TEXT="Content-Range" ID="ID_734624485" CREATED="1590335009364" MODIFIED="1590335013861"/>
<node TEXT="Content-Type" ID="ID_1267748707" CREATED="1590334949366" MODIFIED="1590335152776">
<font BOLD="true"/>
</node>
<node TEXT="Message-Type" ID="ID_864559231" CREATED="1590334953723" MODIFIED="1590335151790">
<font BOLD="true"/>
</node>
<node TEXT="Date" ID="ID_819070803" CREATED="1590335032833" MODIFIED="1590335034264"/>
<node TEXT="ETag" ID="ID_695339938" CREATED="1590335050494" MODIFIED="1590335051677"/>
<node TEXT="Expires" ID="ID_199773883" CREATED="1590335057609" MODIFIED="1590335058967"/>
<node TEXT="Last-Modified" ID="ID_484307697" CREATED="1590335067851" MODIFIED="1590335150551">
<font BOLD="true"/>
</node>
<node TEXT="Link" ID="ID_1425671087" CREATED="1590335134552" MODIFIED="1590335140639"/>
<node TEXT="Location" ID="ID_1526598628" CREATED="1590335137080" MODIFIED="1590335138532"/>
<node TEXT="Retry-After" ID="ID_1750849772" CREATED="1590335141710" MODIFIED="1590335145166">
<font BOLD="true"/>
</node>
</node>
<node TEXT="inbound" ID="ID_1125196905" CREATED="1590334824388" MODIFIED="1590334835169">
<node TEXT="all HTTP request headers may count" ID="ID_463831430" CREATED="1590337336049" MODIFIED="1590337361570"/>
</node>
<node TEXT="outbound" ID="ID_1482796197" CREATED="1590334752714" MODIFIED="1590334839264">
<node TEXT="Allow" ID="ID_1682659937" CREATED="1590334756922" MODIFIED="1590334760808">
<node TEXT="list of Service operations allowed to use (provided by this node)" ID="ID_122085209" CREATED="1590334761239" MODIFIED="1590334781143"/>
<node TEXT="control" ID="ID_1549860379" CREATED="1590334864235" MODIFIED="1590334866382"/>
</node>
</node>
</node>
</node>
</node>
</node>
<node TEXT="body" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_167580041" CREATED="1590331410931" MODIFIED="1590911707724">
<node TEXT="// content" ID="ID_950111002" CREATED="1590331419635" MODIFIED="1590331422379"/>
<node TEXT="mandatory" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1069439007" CREATED="1590911676365" MODIFIED="1590911679673"/>
<node TEXT="default:" ID="ID_206699812" CREATED="1590911760428" MODIFIED="1590911762078">
<node TEXT="{}" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1048648699" CREATED="1590911762510" MODIFIED="1590911764768"/>
</node>
</node>
</node>
<node TEXT="--encoding (?)" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1751128902" CREATED="1590911958083" MODIFIED="1590911966021">
<node TEXT="UTF-8" ID="ID_852246966" CREATED="1590911966430" MODIFIED="1590912079265">
<node TEXT="default" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1452591448" CREATED="1590912057892" MODIFIED="1590912061149"/>
</node>
<node TEXT="protobuf" ID="ID_416792371" CREATED="1590911988274" MODIFIED="1590911994626"/>
<node TEXT="avro" ID="ID_762860769" CREATED="1590911995876" MODIFIED="1590911997946"/>
<node TEXT="ROS (?)" ID="ID_337304087" CREATED="1590911998399" MODIFIED="1590912003458"/>
</node>
<node TEXT="--accept (?)" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_1577474677" CREATED="1590911795729" MODIFIED="1590911920647">
<node TEXT="default:" ID="ID_302066057" CREATED="1590911806927" MODIFIED="1590911809463">
<node TEXT="axon:any" LOCALIZED_STYLE_REF="AutomaticLayout.level,1" ID="ID_851206714" CREATED="1590911809835" MODIFIED="1590911818295"/>
</node>
</node>
</node>
<node TEXT="context" ID="ID_1906182275" CREATED="1590911539694" MODIFIED="1590911541619">
<node TEXT="&#xc1;llapot t&#xe1;rol&#xe1;s&#xe1;ra (is) lehet haszn&#xe1;lni, az infrastrukt&#xfa;ra mellett" ID="ID_771522758" CREATED="1590911542170" MODIFIED="1590911570835"/>
</node>
<node TEXT="channel-ek" ID="ID_1113418893" CREATED="1590331476674" MODIFIED="1590331479576">
<node TEXT="name" ID="ID_21935134" CREATED="1590331515389" MODIFIED="1590331516417">
<node TEXT="symbolic (functional) name" ID="ID_1986363159" CREATED="1590331292083" MODIFIED="1590331512311"/>
<node TEXT="namespace" ID="ID_1844311746" CREATED="1590331299219" MODIFIED="1590331301191"/>
<node TEXT="remapping" ID="ID_227704336" CREATED="1590331296612" MODIFIED="1590331298967"/>
</node>
</node>
<node TEXT="service discovery" ID="ID_201065228" CREATED="1590331763491" MODIFIED="1590331766933">
<node TEXT="orchesctrator channel" ID="ID_1648747025" CREATED="1590331838083" MODIFIED="1605977475884"/>
</node>
<node TEXT="logging" FOLDED="true" ID="ID_327971549" CREATED="1590331849011" MODIFIED="1590339684303">
<node TEXT="lehet t&#xf6;bbf&#xe9;le megold&#xe1;s" ID="ID_1033077446" CREATED="1590339684822" MODIFIED="1590341868519">
<node TEXT="log aggregator, collector" ID="ID_567171684" CREATED="1590341868875" MODIFIED="1590341879394"/>
<node TEXT="dedik&#xe1;lt csatorna" ID="ID_781533696" CREATED="1590341879727" MODIFIED="1590341884674"/>
</node>
</node>
<node TEXT="metrics" FOLDED="true" ID="ID_806077062" CREATED="1590331851341" MODIFIED="1590331853040">
<node TEXT="open-metrics" ID="ID_193441747" CREATED="1590341893616" MODIFIED="1590341900330"/>
<node TEXT="dedik&#xe1;lt csatorna" ID="ID_96291708" CREATED="1590341900606" MODIFIED="1590341904029"/>
</node>
<node TEXT="parameter-server" ID="ID_1187381825" CREATED="1590341904844" MODIFIED="1590341910389">
<node TEXT="kell egy&#xe1;ltal&#xe1;n?" ID="ID_1710362537" CREATED="1590341910716" MODIFIED="1590341915850"/>
<node TEXT="esetleg, ig&#xe9;ny szerint, alkalmaz&#xe1;sf&#xfc;gg&#x151;, k&#xf6;zponti parameter management, t&#xf6;bbf&#xe9;le be&#xe9;ptett szolg&#xe1;ltat&#xe1;ssal." ID="ID_1803148861" CREATED="1590341916206" MODIFIED="1590341955214"/>
<node TEXT="helyette lehetne egy service, ami egy knowledge-graph-b&#xf3;l tud lek&#xe9;rdezni tetsz&#x151;leges dolgot, ak&#xe1;r param&#xe9;tereket" ID="ID_1958730217" CREATED="1605977906251" MODIFIED="1605977934633"/>
</node>
<node TEXT="orchestrator" ID="ID_1621068379" CREATED="1605984475298" MODIFIED="1605984477694">
<node TEXT="A roscore-hoz hasonl&#xf3; funkci&#xf3;kat l&#xe1;t el." ID="ID_819124609" CREATED="1605984478111" MODIFIED="1605984489184"/>
<node TEXT="presence &#xe9;s healthcheck" ID="ID_85392221" CREATED="1605984489551" MODIFIED="1605984516944"/>
<node TEXT="node-okr&#xf3;l r&#xe9;szletes inform&#xe1;ci&#xf3;t tud k&#xe9;rni, &#xe9;s adni" ID="ID_1162487746" CREATED="1605984522816" MODIFIED="1605984538609"/>
<node TEXT="&#xfc;temez&#xe9;st v&#xe9;gez a real-time, szinkron &#xfc;zenetk&#xfc;ld&#xe9;s, feldolgoz&#xe1;si ciklusok &#xf6;sszehangol&#xe1;s&#xe1;hoz." ID="ID_420806540" CREATED="1605984540543" MODIFIED="1605984616730"/>
</node>
<node TEXT="bags" FOLDED="true" ID="ID_542310988" CREATED="1590388246801" MODIFIED="1590388250678">
<node TEXT="recording" ID="ID_12584282" CREATED="1590388251315" MODIFIED="1590388254122">
<node TEXT="with tap" ID="ID_1396073831" CREATED="1590388296563" MODIFIED="1590388300166"/>
</node>
<node TEXT="playing" ID="ID_153197625" CREATED="1590388254353" MODIFIED="1590388257777"/>
<node TEXT="combine maestro with playing, incl. templating and generator" ID="ID_172785858" CREATED="1590388258018" MODIFIED="1590388272291"/>
</node>
</node>
</node>
<node TEXT="glossary:" POSITION="right" ID="ID_651809646" CREATED="1605984859996" MODIFIED="1605984862330">
<node TEXT="RTF" FOLDED="true" ID="ID_176189038" CREATED="1594879420846" MODIFIED="1594879424381">
<node TEXT="Real-Time Factor" ID="ID_600586515" CREATED="1594879427518" MODIFIED="1594879432478"/>
<node TEXT="&quot;Real time&quot; refers to the actual time that is passing in real life as the simulator runs." ID="ID_395542448" CREATED="1594879425663" MODIFIED="1594879453284"/>
<node TEXT="The relationship between the simulation time and real time is known as the &quot;real time factor&quot; (RTF)." ID="ID_414962148" CREATED="1594879454434" MODIFIED="1594879454434"/>
<node TEXT="It&apos;s the ratio of simulation time to real time." ID="ID_1849828993" CREATED="1594879454434" MODIFIED="1594879454434"/>
<node TEXT="The RTF is a measure of how fast or slow your simulation is running compared to real time." ID="ID_1056902048" CREATED="1594879454437" MODIFIED="1594879454437"/>
</node>
</node>
<node TEXT="references:" POSITION="right" ID="ID_1662048566" CREATED="1458311898508" MODIFIED="1458316366680">
<node TEXT="axon (compoNet.mm)" ID="ID_1992629414" CREATED="1605977783623" MODIFIED="1605977817328" LINK="../../axon/docs/compoNet.mm"/>
<node TEXT="R2G1" ID="ID_1827681871" CREATED="1605977640778" MODIFIED="1605977662452" LINK="../../r2g1/docs/r2g1.mm"/>
<node TEXT="ROS.mm" ID="ID_845776089" CREATED="1605977599446" MODIFIED="1605977632195" LINK="../../../Documents/library/Books/AI,%20Brain,%20Robotics/ROS/ROS.mm"/>
</node>
</node>
</map>
