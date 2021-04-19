package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"regexp"
	"strings"
	"testing"
)

func trimSpace(_s string) string {
	s := strings.ReplaceAll(_s, "\u00A0", "\u0020")
	text := strings.TrimSpace(s)
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(text, "")
}

func TestTidy(t *testing.T) {

	html := `
    <div class="top-tool ">
<a class="add-sub-icon top-tool-icon" href="javascript:;" title="添加义项" nslog-type="50000101">
<em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_add-subLemma-solid"></em>
</a>
<a href="/divideload/%E7%B2%9F%E8%A3%95" title="拆分词条" target="_blank" class="split-icon top-tool-icon" style="display:none;" nslog-type="50000104">
<em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_lemma-split"></em>
</a>
<div class="top-collect top-tool-icon" nslog="area" nslog-type="50000102">
<em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_star-solid"></em>
<span class="collect-text">收藏</span>
<div class="collect-tip">查看<a href="/uc/favolemma" target="_blank">我的收藏</a></div>
</div>
<a href="javascript:void(0);" id="j-top-vote" class="top-vote top-tool-icon" nslog-type="10060801">
<em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_zan-solid"></em>
<span class="vote-count">0</span>
<span class="vote-tip">有用+1</span>
<span class="voted-tip">已投票</span>
</a><div class="bksharebuttonbox top-share">
<a class="top-share-icon top-tool-icon" nslog-type="9067">
<em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_share"></em>
<span class="share-count" id="j-topShareCount">0</span>
</a>
<div class="new-top-share" id="top-share">
<ul class="top-share-list">
<li class="top-share-item">
<a class="share-link bds_qzone" href="javascript:void(0);" nslog-type="10060501">
<em class="cmn-icon cmn-icons cmn-icons_logo-qzone"></em>
</a>
</li>
<li class="top-share-item">
<a class="share-link bds_tsina" href="javascript:void(0);" nslog-type="10060701">
<em class="cmn-icon cmn-icons cmn-icons_logo-sina-weibo"></em>
</a>
</li>
<li class="top-share-item">
<a class="bds_wechat" href="javascript:void(0);" nslog-type="10060401">
<em class="cmn-icon cmn-icons cmn-icons_logo-wechat"></em>
</a>
</li>
<li class="top-share-item">
<a class="share-link bds_tqq" href="javascript:void(0);" nslog-type="10060601">
<em class="cmn-icon cmn-icons cmn-icons_logo-qq"></em>
</a>
</li>
</ul>
</div>
</div>
</div>
<div style="width:0;height:0;clear:both"></div><dl class="lemmaWgt-lemmaTitle lemmaWgt-lemmaTitle-">
<dd class="lemmaWgt-lemmaTitle-title">
<h1>粟裕</h1>
<a href="javascript:;" class="edit-lemma cmn-btn-hover-blue cmn-btn-28 j-edit-link"><em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_edit-lemma"></em>编辑</a>
<a class="lock-lemma" nslog-type="10003105" target="_blank" href="/view/10812319.htm" title="锁定"><em class="cmn-icon wiki-lemma-icons wiki-lemma-icons_lock-lemma"></em>锁定</a>
</dd>
</dl><div class="promotion-declaration">
</div><div class="lemma-summary" label-module="lemmaSummary">
<div class="para" label-module="para">粟裕<i>（1907年8月10日-1984年2月5日）</i>，原名粟多珍 ，曾用名粟志裕，<a target="_blank" href="/item/%E4%BE%97%E6%97%8F/154115" data-lemmaid="154115">侗族</a>，生于湖南会同。中国无产阶级革命家、军事家<sup class="sup--normal" data-sup="1" data-ctrmap=":1,">
[1]</sup><a class="sup-anchor" name="ref_[1]_1827"> </a>
，中国人民解放军的主要领导人，中华人民共和国十大<a target="_blank" href="/item/%E5%A4%A7%E5%B0%86/3065" data-lemmaid="3065">大将</a>之首。<sup class="sup--normal" data-sup="1" data-ctrmap=":1,">
[1]</sup><a class="sup-anchor" name="ref_[1]_1827"> </a>
</div><div class="para" label-module="para">1927年加入<a target="_blank" href="/item/%E4%B8%AD%E5%9B%BD%E5%85%B1%E4%BA%A7%E5%85%9A/117227" data-lemmaid="117227">中国共产党</a>，参加<a target="_blank" href="/item/%E5%8D%97%E6%98%8C%E8%B5%B7%E4%B9%89/292100" data-lemmaid="292100">南昌起义</a>，后进入<a target="_blank" href="/item/%E4%BA%95%E5%86%88%E5%B1%B1/28334" data-lemmaid="28334">井冈山</a>，参加历次反“会剿”和全部五次反“围剿”战争。<a target="_blank" href="/item/%E9%95%BF%E5%BE%81">长征</a>时留在南方组织游击战争。抗日战争期间，任<a target="_blank" href="/item/%E6%96%B0%E5%9B%9B%E5%86%9B">新四军</a>第二支队副司令员、江南指挥部和苏北指 挥部副指挥。1941年任<a target="_blank" href="/item/%E6%96%B0%E5%9B%9B%E5%86%9B%E7%AC%AC%E4%B8%80%E5%B8%88/10159917" data-lemmaid="10159917">新四军第一师</a><a target="_blank" href="/item/%E5%B8%88%E9%95%BF/4152239" data-lemmaid="4152239">师长</a>，后兼第六师师长。解放战争期间，任<a target="_blank" href="/item/%E5%8D%8E%E4%B8%AD%E9%87%8E%E6%88%98%E5%86%9B/5695535" data-lemmaid="5695535">华中野战军</a><a target="_blank" href="/item/%E5%8F%B8%E4%BB%A4/4158167" data-lemmaid="4158167">司令</a>、<a target="_blank" href="/item/%E5%8D%8E%E4%B8%9C%E9%87%8E%E6%88%98%E5%86%9B/523256" data-lemmaid="523256">华东野战军</a>副司令、代<a target="_blank" href="/item/%E5%8F%B8%E4%BB%A4%E5%91%98/3825354" data-lemmaid="3825354">司令员</a>兼代<a target="_blank" href="/item/%E6%94%BF%E5%A7%94">政委</a>等职，主要指挥<a target="_blank" href="/item/%E9%AB%98%E9%82%AE%E6%88%98%E5%BD%B9/8630889" data-lemmaid="8630889">高邮战役</a>、<a target="_blank" href="/item/%E9%99%87%E6%B5%B7%E7%BA%BF">陇海线</a>徐（州）海（州）段战役、<a target="_blank" href="/item/%E8%8B%8F%E4%B8%AD%E6%88%98%E5%BD%B9/9192210" data-lemmaid="9192210">苏中战役</a>、<a target="_blank" href="/item/%E5%AD%9F%E8%89%AF%E5%B4%AE%E6%88%98%E5%BD%B9/2903825" data-lemmaid="2903825">孟良崮战役</a>、<a target="_blank" href="/item/%E6%B5%8E%E5%8D%97%E6%88%98%E5%BD%B9/18093" data-lemmaid="18093">济南战役</a>、<a target="_blank" href="/item/%E6%B7%AE%E6%B5%B7%E6%88%98%E5%BD%B9/287777" data-lemmaid="287777">淮海战役</a> 、<a target="_blank" href="/item/%E6%B8%A1%E6%B1%9F%E6%88%98%E5%BD%B9/283718" data-lemmaid="283718">渡江战役</a>、<a target="_blank" href="/item/%E4%B8%8A%E6%B5%B7%E6%88%98%E5%BD%B9/10520224" data-lemmaid="10520224">上海战役</a>等。中华人民共和国成立后，历任中国人民解放军<a target="_blank" href="/item/%E6%80%BB%E5%8F%82%E8%B0%8B%E9%95%BF/6026043" data-lemmaid="6026043">总参谋长</a>、<a target="_blank" href="/item/%E4%B8%AD%E5%9B%BD%E5%85%B1%E4%BA%A7%E5%85%9A%E4%B8%AD%E5%A4%AE%E5%86%9B%E4%BA%8B%E5%A7%94%E5%91%98%E4%BC%9A/2282542" data-lemmaid="2282542">中国共产党中央军事委员会</a>常委、第五届全国人大常委会副委员长等职。</div><div class="para" label-module="para">1955年9月27日，被授予<a target="_blank" href="/item/%E5%A4%A7%E5%B0%86/3065" data-lemmaid="3065">大将</a>军衔，并授予一级<a target="_blank" href="/item/%E5%85%AB%E4%B8%80%E5%8B%8B%E7%AB%A0/8462555" data-lemmaid="8462555">八一勋章</a>、一级<a target="_blank" href="/item/%E7%8B%AC%E7%AB%8B%E8%87%AA%E7%94%B1%E5%8B%8B%E7%AB%A0/8462589" data-lemmaid="8462589">独立自由勋章</a>和一级<a target="_blank" href="/item/%E8%A7%A3%E6%94%BE%E5%8B%8B%E7%AB%A0/8462624" data-lemmaid="8462624">解放勋章</a>。</div>
</div>
<div class="lemmaWgt-promotion-leadPVBtn">
</div><div class="configModuleBanner">
</div><div class="lemmaWgt-focusAndRelation relations">
<div class="lemma-relation-title"><i>人物关系</i></div>
<div id="fixBox" style="right: 10px;">
<div class="openFix" title="若关系内容有误，请点击纠错进行反馈">纠错</div>
<div class="closeFix">关闭纠错</div>
</div>
<div class="lemma-relation-module viewport" id="J-lemma-relation-module">
<ul class="lemma-relation-list slider maqueeCanvas">
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E6%A5%9A%E9%9D%92/10356628" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/9825bc315c6034a8ba70233ac1134954082376ff-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="楚青">
<span class="name">妻子</span>
<span class="title">楚青</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">10356628</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E7%B2%9F%E5%AF%92%E7%94%9F/8809376" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/4610b912c8fcc3ce20ef21aa9845d688d43f2078-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="粟寒生">
<span class="name">儿子</span>
<span class="title">粟寒生</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">8809376</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E7%B2%9F%E6%88%8E%E7%94%9F/8521657" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/71cf3bc79f3df8dc79f5ada7ce11728b4710283f-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="粟戎生">
<span class="name">儿子</span>
<span class="title">粟戎生</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">8521657</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E5%88%98%E4%BA%9A%E6%A5%BC/1202954" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/0ff41bd5ad6eddc451daab84ad91a1fd5266d0169530-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="刘亚楼">
<span class="name">战友</span>
<span class="title">刘亚楼</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">1202954</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E6%9C%B1%E5%BE%B7/115094" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/58c3acb758b6df9431add1d2-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="朱德">
<span class="name">战友</span>
<span class="title">朱德</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">115094</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E9%99%88%E6%AF%85/22586" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/95eef01f3a292df5f2101b00bf315c6035a873d7-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="陈毅">
<span class="name">战友</span>
<span class="title">陈毅</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">22586</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
<li class="lemma-relation-item">
<a class="lemma-relation-link" href="/item/%E5%88%98%E4%BC%AF%E6%89%BF/116201" nslog-type="10000302" target="_blank">
<img src="https://bkimg.cdn.bcebos.com/smart/4b90f603738da97753b03d24bf51f8198618e326-bkimg-process,v_1,rw_1,rh_1,maxl_400?x-bce-process=image/format,f_auto"/>
<div class="info" title="刘伯承">
<span class="name">战友</span>
<span class="title">刘伯承</span>
</div>
</a>
<div class="hasErrorBox hasErrorHidden">
<a class="hasError" href="">
有错误<span class="relateId" style="display:none;">116201</span></a>
<span class="hadSubError">已反馈</span>
</div>
</li>
</ul>
</div>
<form id="errFormSub" action="https://sp2.baidu.com/-uV1bjeh1BF3odCf/index.php/feedback/zx/baikeJC" target="userSubErr" method="post">
<input id="source_id" name="source_id" type="text" value=""/>
<input id="errForm_query" name="query" type="text" value=""/>
<input id="html_url" name="html_url" type="text" value=""/>
<input id="img_url" name="img_url" type="text" value=""/>
</form>
<iframe id="userSubErr" name="userSubErr" style="display:none;" frameborder="0">
</iframe>
</div>
<div class="basic-info cmn-clearfix">
<dl class="basicInfo-block basicInfo-left">
<dt class="basicInfo-item name">中文名</dt>
<dd class="basicInfo-item value">
粟裕
</dd>
<dt class="basicInfo-item name">外文名</dt>
<dd class="basicInfo-item value">
Su Yu
</dd>
<dt class="basicInfo-item name">别    名</dt>
<dd class="basicInfo-item value">
粟多珍、粟志裕
</dd>
<dt class="basicInfo-item name">国    籍</dt>
<dd class="basicInfo-item value">
<a target="_blank" href="/item/%E4%B8%AD%E5%9B%BD/22516505" data-lemmaid="22516505">中国</a>
</dd>
<dt class="basicInfo-item name">民    族</dt>
<dd class="basicInfo-item value">
<a target="_blank" href="/item/%E4%BE%97%E6%97%8F">侗族</a>
</dd>
<dt class="basicInfo-item name">出生日期</dt>
<dd class="basicInfo-item value">
1907年8月10日
</dd>
<dt class="basicInfo-item name">逝世日期</dt>
<dd class="basicInfo-item value">
1984年2月5日
</dd>
</dl><dl class="basicInfo-block basicInfo-right">
<dt class="basicInfo-item name">毕业院校</dt>
<dd class="basicInfo-item value">
湖南省立第二师范
</dd>
<dt class="basicInfo-item name">职    业</dt>
<dd class="basicInfo-item value">
军事家、革命家
</dd>
<dt class="basicInfo-item name">主要成就</dt>
<dd class="basicInfo-item value">
抗日先遣、苏中抗战、华东主将、国防（建设
<br/>车桥、孟良崮、豫东、淮海）等战役主要指挥
</dd>
<dt class="basicInfo-item name">出生地</dt>
<dd class="basicInfo-item value">
湖南会同
</dd>
<dt class="basicInfo-item name">代表作品</dt>
<dd class="basicInfo-item value">
《粟裕战争回忆录》、《实战经验录》、《粟裕文选》
</dd>
<dt class="basicInfo-item name">军    衔</dt>
<dd class="basicInfo-item value">
大将
</dd>
</dl></div>
<div class="lemmaWgt-lemmaCatalog">
<div class="lemma-catalog">
<h2 class="block-title">目录</h2>
<div class="catalog-list column-4">
<ol>
<li class="level1">
<span class="index">1</span>
<span class="text"><a href="#1">人物生平</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#1_1">青年求学</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#1_2">入党革命</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#1_3">抗战时期</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#1_4">解放战争时期</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#1_5">建国初期</a></span>
</li>
</ol><ol><li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#1_6">建国后期</a></span>
</li>
<li class="level1">
<span class="index">2</span>
<span class="text"><a href="#2">代表著作</a></span>
</li>
<li class="level1">
<span class="index">3</span>
<span class="text"><a href="#3">人物成就</a></span>
</li>
<li class="level1">
<span class="index">4</span>
<span class="text"><a href="#4">家庭成员</a></span>
</li>
<li class="level1">
<span class="index">5</span>
<span class="text"><a href="#5">轶事典故</a></span>
</li>
</ol><ol><li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#5_1">夫人赋诗</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#5_2">颅中弹片</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#5_3">授衔让帅</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#5_4">艰难平反</a></span>
</li>
<li class="level1">
<span class="index">6</span>
<span class="text"><a href="#6">纪念场所</a></span>
</li>
<li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#6_1">纪念堂</a></span>
</li>
</ol><ol><li class="level2">
<span class="index">▪</span>
<span class="text"><a href="#6_2">纪念馆</a></span>
</li>
<li class="level1">
<span class="index">7</span>
<span class="text"><a href="#7">人物评价</a></span>
</li>
<li class="level1">
<span class="index">8</span>
<span class="text"><a href="#8">影视形象</a></span>
</li>
</ol>

</div>
</div>
</div>
<div class="anchor-list ">
<a name="1" class="lemma-anchor para-title"></a>
<a name="sub1827_1" class="lemma-anchor "></a>
<a name="人物生平" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>人物生平</h2>
</div>
<div class="anchor-list ">
<a name="1_1" class="lemma-anchor para-title"></a>
<a name="sub1827_1_1" class="lemma-anchor "></a>
<a name="青年求学" class="lemma-anchor "></a>
<a name="1-1" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>青年求学</h3>
</div>
<div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/5327ce165ab0dd3f972b4398?fr=lemma&amp;ct=single" target="_blank" title="青年粟裕" style="width:220px;height:309.85915492958px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/738b4710b912c8fc905ab5dafc039245d688217f?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="青年粟裕" style="width:220px;height:309.85915492958px;"/>
</a>
<span class="description">
青年粟裕
<sup>[2]</sup>
<a name="ref_[2]_"></a>
</span>
</div>1907年8月10日<i>（农历丁未年、清光绪三十三年）</i>，出生于湖南<a target="_blank" href="/item/%E6%80%80%E5%8C%96%E5%B8%82">怀化市</a>会同县伏龙乡（今坪村）枫木树脚村，幼名继业，学名多珍，字裕。</div><div class="para" label-module="para">1916年转入其叔父创办的第八国民学校读书。</div><div class="para" label-module="para">1918年，迁居<a target="_blank" href="/item/%E4%BC%9A%E5%90%8C%E5%8E%BF">会同县</a>城，先后入粟氏私立初级国民学校和会同县立第一高等小学读书。</div><div class="para" label-module="para">1923年报考湖南省立第二师范，被录取为选送生。</div><div class="para" label-module="para">1924年3月，因错过考期，进入湖南省立第二师范学 校（常德）附小和平民中学学习。</div><div class="para" label-module="para">1925年春，考上省立二师，后因省立二师进步校长被害，粟裕等进步学生被秘密转移到<a target="_blank" href="/item/%E6%AD%A6%E6%98%8C">武昌</a>，安排到<a target="_blank" href="/item/%E5%8F%B6%E6%8C%BA">叶挺</a>24师教导大队，任学员班长。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="anchor-list ">
<a name="1_2" class="lemma-anchor para-title"></a>
<a name="sub1827_1_2" class="lemma-anchor "></a>
<a name="入党革命" class="lemma-anchor "></a>
<a name="1-2" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>入党革命</h3>
</div>
<div class="para" label-module="para">1926年11月，粟裕加入<a target="_blank" href="/item/%E4%B8%AD%E5%9B%BD%E5%85%B1%E4%BA%A7%E4%B8%BB%E4%B9%89%E9%9D%92%E5%B9%B4%E5%9B%A2">中国共产主义青年团</a>。1927年6月，粟裕转入中国共产党。</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/63d9f2d3572c11df7474cc9b662762d0f603c2e5?fr=lemma&amp;ct=single" target="_blank" title="粟裕" style="width:220px;height:347.82608695652px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/63d9f2d3572c11df7474cc9b662762d0f603c2e5?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕" style="width:220px;height:347.82608695652px;"/>
</a>
<span class="description">
粟裕
</span>
</div>1927年8月1日，他参加了著名的<a target="_blank" href="/item/%E5%8D%97%E6%98%8C%E8%B5%B7%E4%B9%89">南昌起义</a>，任起义军总指挥部警卫队班长。</div><div class="para" label-module="para">1928年1月，粟裕参加<a target="_blank" href="/item/%E6%B9%98%E5%8D%97%E8%B5%B7%E4%B9%89">湘南起义</a>后到了<a target="_blank" href="/item/%E4%BA%95%E5%86%88%E5%B1%B1">井冈山</a>。<sup class="sup--normal" data-sup="3" data-ctrmap=":3,">
[3]</sup><a class="sup-anchor" name="ref_[3]_1827"> </a>
</div><div class="para" label-module="para">1929年后，因屡立战功粟裕相继升任<a target="_blank" href="/item/%E8%90%A5%E9%95%BF">营长</a>、团长、<a target="_blank" href="/item/%E5%B8%88%E9%95%BF">师长</a>，红四军参谋长，红十一军参谋长，<a target="_blank" href="/item/%E7%BA%A2%E4%B8%83%E5%86%9B%E5%9B%A2">红七军团</a>参谋长等职。</div><div class="para" label-module="para">1932年2月，粟裕由红军学校调回红四军，仍任红四军参谋长。12月，任红一军团教导师政委兼政治部主任。</div><div class="para" label-module="para">1934年11月，粟裕调任闽浙赣军区参谋长。根据中共中央革命军事委员会命令，红军北上抗日先遣队与方志敏领导的红十军及地方武装合编，成立红军第十军团。下旬，红十军团奉命转到<a target="_blank" href="/item/%E5%A4%96%E7%BA%BF%E4%BD%9C%E6%88%98">外线作战</a>，调任红十军团参谋长。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="para" label-module="para">1月红十军团在谭家桥战斗失利后，又在怀玉山遭到敌 人围歼。率领先头部队果断突出封锁线，安全到达闽浙赣苏区。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
2月-4月，建立浙西南游击区，任挺进师（四百余人）师长。5月-8月，与<a target="_blank" href="/item/%E5%88%98%E8%8B%B1/1605" data-lemmaid="1605">刘英</a>一起指挥粉碎国民党军队对挺进师的第一次“进剿”，挺进师发展到近千人。10月5日，中共闽浙边临时省委和闽浙边临时省军区成立，任省军区司 令员、省委组织部长。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="anchor-list ">
<a name="1_3" class="lemma-anchor para-title"></a>
<a name="sub1827_1_3" class="lemma-anchor "></a>
<a name="抗战时期" class="lemma-anchor "></a>
<a name="1-3" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>抗战时期</h3>
</div>
<div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/fcfaaf51f3deb48fac2efa25f41f3a292cf57881?fr=lemma&amp;ct=single" target="_blank" title="粟裕" style="width:220px;height:301.08571428571px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/fcfaaf51f3deb48fac2efa25f41f3a292cf57881?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕" style="width:220px;height:301.08571428571px;"/>
</a>
<span class="description">
粟裕
</span>
</div>1938年3月18日，粟裕率浙闽边抗日游击总队从<a target="_blank" href="/item/%E5%B9%B3%E9%98%B3%E5%8E%BF">平阳县</a>山门街开赴皖南，加入新四军战斗行列。部队整编为新四军第二支队第四团 第三营，任第二支队副司令员。4月28日，他奉命组建新四军先遣支队，任先遣支队司令员，向<a target="_blank" href="/item/%E8%8B%8F%E5%8D%97">苏南</a>敌后执行侦察任务。6月11日，<a target="_blank" href="/item/%E5%A5%89%E5%91%BD">奉命</a>执行挺进南京、镇江间破坏铁道任务。6月17日，在韦 岗伏击日军，歼灭日军<a target="_blank" href="/item/%E5%B0%91%E4%BD%90">少佐</a>土井以下官兵30多人。6月21日，先遣支队撤销，仍回第二支队任副司令员，后任代司令员。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1939年1月，粟裕在指挥<a target="_blank" href="/item/%E6%B0%B4%E9%98%B3%E9%95%87">水阳镇</a>伏击战、<a target="_blank" href="/item/%E6%A8%AA%E5%B1%B1/63986" data-lemmaid="63986">横山</a>战斗、<a target="_blank" href="/item/%E5%A5%87%E8%A2%AD%E5%AE%98%E9%99%A1%E9%97%A8">奇袭官陡门</a>等战斗中，歼日伪军400余人，俘日伪军57名，并炸毁火车一列。8月，新四军江南指挥部成立，任副指挥。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/b258f5c468b44aee8326ac1e?fr=lemma&amp;ct=single" target="_blank" title="粟裕与夫人楚青" style="width:220px;height:293.33333333333px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/9f510fb30f2442a706c2c1d2d143ad4bd01302fc?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕与夫人楚青" style="width:220px;height:293.33333333333px;"/>
</a>
<span class="description">
粟裕与夫人楚青
</span>
</div>1941年1月“<a target="_blank" href="/item/%E7%9A%96%E5%8D%97%E4%BA%8B%E5%8F%98">皖南事 变</a>”后，粟裕任<a target="_blank" href="/item/%E6%96%B0%E5%9B%9B%E5%86%9B%E7%AC%AC%E4%B8%80%E5%B8%88">新四军第一师</a>师长（后兼政治委员），苏中军区司令员兼政治委员。8月13日，指挥苏 中军民反击日伪军报复性的“扫荡”，连续作战42昼夜、130余次，歼日军1300余人。8月中旬起，领导和指挥持续8个月的要点争夺战，“七保三仓”，“五保丰利”（毙伤日军800多人，伪军更多<sup class="sup--normal" data-sup="5" data-ctrmap=":5,">
[5]</sup><a class="sup-anchor" name="ref_[5]_1827"> </a>
），保持了相对稳定的根据地基本区。12月26日，与<a target="_blank" href="/item/%E6%A5%9A%E9%9D%92">楚青</a>在黄海之滨石家庄（今江苏省如皋市<a target="_blank" href="/item/%E7%9F%B3%E5%BA%84%E9%95%87">石庄镇</a>）结婚。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1944年1、2月间，粟裕发起春季攻势作战，解放国土近三千平方公里、村镇一百五十多处，争取日伪军一千余人反正。3月，组织指挥<a target="_blank" href="/item/%E8%BD%A6%E6%A1%A5%E6%88%98%E5%BD%B9">车桥战役</a>，歼日军三泽大佐以下官兵460余人、 伪军480余人，摧毁日军碉堡50座。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
6月26日，发起<a target="_blank" href="/item/%E5%8D%97%E5%9D%8E%E6%88%98%E5%BD%B9">南坎战役</a>，共拔除日伪据点七八十处。9月21日-10月31日，组织指挥讨陈战役，歼灭<a target="_blank" href="/item/%E9%99%88%E6%B3%B0%E8%BF%90">陈泰运</a>部及日伪军2300余人。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="para" label-module="para">1945年1月13日，粟裕任苏浙军区司令员，统一指挥苏 南、浙西、浙东部队。10月任华中军区副司令员、<a target="_blank" href="/item/%E5%8D%8E%E4%B8%AD%E9%87%8E%E6%88%98%E5%86%9B">华中野战军</a>司令员，指挥<a target="_blank" href="/item/%E9%AB%98%E9%82%AE%E6%88%98%E5%BD%B9">高邮战役</a>和陇海线徐（州）海（州）段战役，歼灭拒降日伪军2万余人，为迎击国民党军的进攻准备了内线作战的有利条件，使华中、山东解放区连成一片。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="anchor-list ">
<a name="1_4" class="lemma-anchor para-title"></a>
<a name="sub1827_1_4" class="lemma-anchor "></a>
<a name="解放战争时期" class="lemma-anchor "></a>
<a name="1-4" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>解放战争时期</h3>
</div>
<div class="para" label-module="para">1946年6月<a target="_blank" href="/item/%E8%92%8B%E4%BB%8B%E7%9F%B3">蒋介石</a>发动全面内战后，<a target="_blank" href="/item/%E4%B8%AD%E5%85%B1%E4%B8%AD%E5%A4%AE">中共中央</a>采纳粟裕的建议，改变太行、山东、华中3支大军同时出击外线的计划 ，同意华中野战军主力先在苏中内线作战。</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/7aec54e736d12f2e8c86d3f94ac2d562843568a3?fr=lemma&amp;ct=single" target="_blank" title="粟裕大将" style="width:220px;height:280.04474272931px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/7aec54e736d12f2e8c86d3f94ac2d562843568a3?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕大将" style="width:220px;height:280.04474272931px;"/>
</a>
<span class="description">
粟裕大将
</span>
</div>1946年7月中旬，国民党集中正规军50万人，向华东解放区的华中野战军聚集地发起进攻。粟裕、<a target="_blank" href="/item/%E8%B0%AD%E9%9C%87%E6%9E%97">谭震林</a>指挥华中野战军19个团3万余人奋起迎击。分两个阶段作战，历时45天，歼灭国民党军6个旅、5个交通警察大队共5.3万人（国民党 军五分之二兵力），取得了在内线歼灭美械装备的国民党军的初步经验，是人民解放军在解放战争初期取得的重大胜利之一。<sup class="sup--normal" data-sup="6" data-ctrmap=":6,">
[6]</sup><a class="sup-anchor" name="ref_[6]_1827"> </a>
</div><div class="para" label-module="para">1947年1月，粟裕率华东野战军先后发起了<a target="_blank" href="/item/%E5%AE%BF%E5%8C%97%E6%88%98%E5%BD%B9">宿北战役</a>、<a target="_blank" href="/item/%E9%B2%81%E5%8D%97%E6%88%98%E5%BD%B9">鲁南战役</a>、<a target="_blank" href="/item/%E8%8E%B1%E8%8A%9C%E6%88%98%E5%BD%B9">莱芜战役</a>、<a target="_blank" href="/item/%E6%B3%B0%E8%92%99%E6%88%98%E5%BD%B9">泰蒙战役</a>、<a target="_blank" href="/item/%E5%AD%9F%E8%89%AF%E5%B4%AE%E6%88%98%E5%BD%B9">孟良崮战役</a>等，共歼国民党7个军（整编师）和1个快速纵队。其 中，1947年5月歼灭国民党号称“王牌军”的整编第<a target="_blank" href="/item/74%E5%B8%88">74师</a>。人民解放军转入战略进攻后，率<a target="_blank" href="/item/%E5%8D%8E%E4%B8%9C%E9%87%8E%E6%88%98%E5%86%9B">华东野战军</a>主力挺进<a target="_blank" href="/item/%E9%B2%81%E8%A5%BF%E5%8D%97">鲁西南</a>，掩护<a target="_blank" href="/item/%E6%99%8B%E5%86%80%E9%B2%81%E8%B1%AB%E9%87%8E%E6%88%98%E5%86%9B">晋冀鲁豫野战军</a>主力南下<a target="_blank" href="/item/%E5%A4%A7%E5%88%AB%E5%B1%B1">大别山</a>，指挥<a target="_blank" href="/item/%E6%B2%99%E5%9C%9F%E9%9B%86%E6%88%98%E5%BD%B9">沙土集战役</a>，歼国民党整编第57师，迫使国民军从山东和<a target="_blank" href="/item/%E5%A4%A7%E5%88%AB%E5%B1%B1">大别山</a>区抽调4个整编师来援，实现了华东战区由内线向外线、从战略防御到战略进攻的转折，随即挺进<a target="_blank" href="/item/%E8%B1%AB%E7%9A%96%E8%8B%8F%E8%BE%B9%E5%8C%BA">豫皖苏边区</a>。</div><div class="para" label-module="para">1948年5月，被任命为<a target="_blank" href="/item/%E5%8D%8E%E4%B8%9C%E9%87%8E%E6%88%98%E5%86%9B">华东野战军</a>司令兼政委，在其推辞后任代司令兼代政委。1948年6月兼任<a target="_blank" href="/item/%E8%B1%AB%E7%9A%96%E8%8B%8F">豫皖苏</a>军区司令员。</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/787014556ec7e283b645ae03?fr=lemma&amp;ct=single" target="_blank" title="粟裕与陈毅" style="width:220px;height:155.96428571429px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/e7cd7b899e510fb3d2fa4be0d933c895d0430ce1?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕与陈毅" style="width:220px;height:155.96428571429px;"/>
</a>
<span class="description">
粟裕与陈毅
</span>
</div>1948年6月16日，粟裕发起开封战役（又称<a target="_blank" href="/item/%E8%B1%AB%E4%B8%9C%E6%88%98%E5%BD%B9">豫东战役</a>第一阶段），激战5昼夜，22日攻克开封，全歼守敌及部分援敌共4万余人。6月27日发起睢杞战役（又称豫东战役第二阶段），激战6天，共歼国民党军5万余人。7月12日发起<a target="_blank" href="/item/%E5%85%96%E5%B7%9E%E6%88%98%E5%BD%B9">兖州战役</a>，歼敌6.3万余人，使王耀武盘踞的济南陷于孤立。豫东战役中，华东野战军参战兵力达20万人，国民党军参战的兵力达25万余人。经过20天连续作战，华东野战军共歼敌9万余人，给了中原之敌以重创，为华东野战军进一 步开展中原和华东战局，创造了有利条件。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,:7,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
<sup class="sup--normal" data-sup="7" data-ctrmap=":4,:7,">
[7]</sup><a class="sup-anchor" name="ref_[7]_1827"> </a>
</div><div class="para" label-module="para">1948年9月11日，被任命华东野战军总指挥。9月16日，发起<a target="_blank" href="/item/%E6%B5%8E%E5%8D%97%E6%88%98%E5%BD%B9/18093" data-lemmaid="18093">济南战役</a>，24日胜利结束，全歼济南守敌10.4万余人（包括起义两万余人），生俘国民党第二绥靖区司令官王耀武以下将领23名。在阻援打援战场上，国民党援军迟迟不敢北上增援，华东野战军不战而胜。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="para" label-module="para">1948年11月6日，粟裕率华东野战军发起<a target="_blank" href="/item/%E6%B7%AE%E6%B5%B7%E6%88%98%E5%BD%B9">淮海战役</a>。该战役共投入解放军66万，<sup class="sup--normal" data-sup="8" data-ctrmap=":8,">
[8]</sup><a class="sup-anchor" name="ref_[8]_1827"> </a>
地方部队40万。<sup class="sup--normal" data-sup="9" data-ctrmap=":9,">
[9]</sup><a class="sup-anchor" name="ref_[9]_1827"> </a>
在战役中，粟裕指指挥<a target="_blank" href="/item/%E5%8D%8E%E4%B8%9C%E9%87%8E%E6%88%98%E5%86%9B">华东野战军</a>17个纵队作战，歼灭国民党军44万余人，解放军伤亡13万余人。战役过后，毛泽东说：“淮海战役，粟裕同志立了第一功。”</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/a5c27d1ed21b0ef41e836333dcc451da81cb3e77?fr=lemma&amp;ct=single" target="_blank" title="在上海我军入城式 上检阅部队" style="width:220px;height:155.1px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/a5c27d1ed21b0ef41e836333dcc451da81cb3e77?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="在上海我军入城式上检阅部队" style="width:220px;height:155.1px;"/>
</a>
<span class="description">
在上海我军入城式上检阅部队
</span>
</div></div><div class="para" label-module="para">1949年1月，粟裕任<a target="_blank" href="/item/%E7%AC%AC%E4%B8%89%E9%87%8E%E6%88%98%E5%86%9B">第三野战军</a>副司令员兼第二副政治委员 。2月开始筹备第三野战军京沪杭战役（实际分为渡江战役、上海战役等）。4月下旬，他指挥的<a target="_blank" href="/item/%E6%B8%A1%E6%B1%9F%E6%88%98%E5%BD%B9">渡江战役</a>解放了南京、杭州。5月指挥<a target="_blank" href="/item/%E4%B8%8A%E6%B5%B7%E6%88%98%E5%BD%B9">上海战役</a>，在 上海外围歼敌主力8个军。先后兼任上海市军管会副主任、南京市军管会主任、南京市市长、<a target="_blank" href="/item/%E5%8D%8E%E4%B8%9C%E5%86%9B%E6%94%BF%E5%A7%94%E5%91%98%E4%BC%9A">华东军 政委员会</a>副主席。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="para" label-module="para"><a target="_blank" href="/item/1949%E5%B9%B4">1949年</a>5月20日，粟裕为上海军事管制委员会副主任。8月2日，华东海军由粟裕指挥。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1949年9月中旬-10月中旬，粟裕率领第三野战军代表团出席在北京召开的全国政治协商会议、中华人民共和国开国大典和中央军委召开的军事会议。会议中粟裕当选为全国政协委员，被中央人民政府任命为中国人民革命军事委员会委员。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1949年9月30日，党和国家领导人在<a target="_blank" href="/item/%E5%A4%A9%E5%AE%89%E9%97%A8%E5%B9%BF%E5%9C%BA">天安门广场</a>为<a target="_blank" href="/item/%E4%BA%BA%E6%B0%91%E8%8B%B1%E9%9B%84%E7%BA%AA%E5%BF%B5%E7%A2%91">人民英雄纪 念碑</a>举行奠基典礼。<a target="_blank" href="/item/%E6%AF%9B%E6%B3%BD%E4%B8%9C">毛泽东</a>主席是第一个上前铲土的，粟裕作为第三野战军代表团团长兼首席代表，紧随<a target="_blank" href="/item/%E6%9C%B1%E5%BE%B7">朱德</a>、<a target="_blank" href="/item/%E8%B4%BA%E9%BE%99">贺龙</a>之后第四个铲土。</div><div class="anchor-list ">
<a name="1_5" class="lemma-anchor para-title"></a>
<a name="sub1827_1_5" class="lemma-anchor "></a>
<a name="建国初期" class="lemma-anchor "></a>
<a name="1-5" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>建国初期</h3>
</div>
<div class="para" label-module="para">中华人民共和国成立后，粟裕历任解放台湾工作委员会主任、中国人民革命军事委员会副总参谋长、中国人民解放军总参谋长、国防部副部长、军事科学院副院长、第一政治委员，中共中央军委常务委员。是中共第八届至第十一届<a target="_blank" href="/item/%E4%B8%AD%E5%A4%AE%E5%A7%94%E5%91%98">中央委员</a>，第一至三届国防委员会委员，第三至第五届全国人 大常委会副委员长，中共中央顾问委员会常委。</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/024f78f0f736afc35b8cdf08b219ebc4b64512de?fr=lemma&amp;ct=single" target="_blank" title="1955年，粟裕与陈 赓在天安门城楼" style="width:220px;height:299.2px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/024f78f0f736afc35b8cdf08b219ebc4b64512de?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="1955年，粟裕与陈赓在天安门城楼" style="width:220px;height:299.2px;"/>
</a>
<span class="description">
1955年，粟裕与陈赓在天安门城楼
</span>
</div>1950年1月27日出席华东军政委员会成立会议，粟裕就任华东军政委员会副主席。6月上旬出席中共七届三中全会。会议期间，粟裕建议中央军委直接指挥台湾战役。毛泽东宣布：攻台作战仍由粟裕负责。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1950年7月上旬，粟裕为东北边防军司令员兼政治委员 。7月14日，粟裕复发高血压、美尼尔氏综合症，到青岛治疗。8月1日，粟裕病情仍未好转。12月，经中 共中央批准，到苏联疗养。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1951年9月，粟裕从苏联回国，<a target="_blank" href="/item/%E5%91%A8%E6%81%A9%E6%9D%A5">周恩来</a>、<a target="_blank" href="/item/%E6%9C%B1%E5%BE%B7">朱德</a>向他传达中共中央要他到总参工作的决定。11月12日，中央军委任命粟裕为<a target="_blank" href="/item/%E4%B8%AD%E5%A4%AE%E4%BA%BA%E6%B0%91%E6%94%BF%E5%BA%9C">中央人民政府</a>革命军事委员会第二副总参谋长，仍兼华东军区副司令员。12月12日，就任<a target="_blank" href="/item/%E4%BA%BA%E6%B0%91%E9%9D%A9%E5%91%BD%E5%86%9B%E4%BA%8B%E5%A7%94%E5%91%98%E4%BC%9A"> 人民革命军事委员会</a>副总参谋长。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para">1954年10月31日，<a target="_blank" href="/item/%E4%B8%AD%E5%85%B1%E4%B8%AD%E5%A4%AE">中共中央</a>通知：粟裕任总参谋长。</div><div class="para" label-module="para">在总参工作期间，参与<a target="_blank" href="/item/%E6%8A%97%E7%BE%8E%E6%8F%B4%E6%9C%9D%E6%88%98%E4%BA%89">抗美援朝战争</a>指导，提出很多重要建议，提出完整的军队建设计划，实现新中国第一代国防体系建设。</div><div class="para" label-module="para">1955年9 月27日，粟裕参加<a target="_blank" href="/item/%E5%91%A8%E6%81%A9%E6%9D%A5">周恩来</a>总理授予军衔的仪式。周恩来把授予大将军衔命令状第一个授予粟裕，并授予一级<a target="_blank" href="/item/%E5%85%AB%E4%B8%80%E5%8B%8B%E7%AB%A0">八一勋章</a>、一级<a target="_blank" href="/item/%E7%8B%AC%E7%AB%8B%E8%87%AA%E7%94%B1%E5%8B%8B%E7%AB%A0">独立自由勋章</a>和一级<a target="_blank" href="/item/%E8%A7%A3%E6%94%BE%E5%8B%8B%E7%AB%A0">解放勋章</a>。</div><div class="para" label-module="para">1956年11月20日，增补粟裕为中央军委委员。</div><div class="anchor-list ">
<a name="1_6" class="lemma-anchor para-title"></a>
<a name="sub1827_1_6" class="lemma-anchor "></a>
<a name="建国后期" class="lemma-anchor "></a>
<a name="1-6" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>建国后期</h3>
</div>
<div class="para" label-module="para">1958年2月，粟裕参加以<a target="_blank" href="/item/%E5%91%A8%E6%81%A9%E6%9D%A5">周恩来</a>总理为首的政府代表团访问朝鲜，粟此行主要与朝方协商部署 志愿军回国事宜。6月30日晚上，他主持总参谋部第四次部务会议，讨论军委扩大会议提出的有关总参工 作和军队建设方面的重大问题。9月19日，正式到军事科学院任副院长、党委第一副书记，负责院的常务 工作。12月30日，<a target="_blank" href="/item/%E5%9B%BD%E9%98%B2%E9%83%A8">国防部</a>命令：粟裕以国防部副部长职务兼任<a target="_blank" href="/item/%E5%86%9B%E4%BA%8B%E7%A7%91%E5%AD%A6%E9%99%A2">军事科学院</a>副院长。<sup class="sup--normal" data-sup="2" data-ctrmap=":2,">
[2]</sup><a class="sup-anchor" name="ref_[2]_1827"> </a>
</div><div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/b90e7bec54e736d1839a6d659a504fc2d46269d6?fr=lemma&amp;ct=single" target="_blank" title="粟裕一叶帅在军科 院会议上" style="width:220px;height:127.05px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/b90e7bec54e736d1839a6d659a504fc2d46269d6?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕一叶帅在军科院会议上" style="width:220px;height:127.05px;"/>
</a>
<span class="description">
粟裕一叶帅在军科院会议上
</span>
</div></div><div class="para" label-module="para">1970年6月，粟裕率中国军事代表团访问刚果， 途经法国考察了诺曼底。</div><div class="para" label-module="para">1973年8月24日，粟裕出席中 共十大，在再次被选为中央委员。</div><div class="para" label-module="para">1975年1月13日，粟 裕出席第四届全国人民代表大会第一次会议，任解放军代表团团长，被选为人大常务委员会委员。中共中央决定成立中共中央军事委员会常务委员会，他被任命为常务委员会委员。</div><div class="para" label-module="para">1977年8月，粟裕出席中共十一大被选为中央委员。1982年5月7日，应《淮海千秋》 摄制组要求，他参加有关淮海战役的座谈会。同年，被任命为中央顾问委员会常委。</div><div class="para" label-module="para">1984年2月5日16时33分，粟裕在中国人民解放军总医院逝世。<sup class="sup--normal" data-sup="10-11" data-ctrmap=":10,:11,">
[10-11]</sup><a class="sup-anchor" name="ref_[10-11]_1827"> </a>
</div><div class="anchor-list ">
<a name="2" class="lemma-anchor para-title"></a>
<a name="sub1827_2" class="lemma-anchor "></a>
<a name="代表著作" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>代表著作</h2>
</div>
<table log-set-param="table_view" data-sort="sortDisabled"><tbody><tr><td align="middle" width="170"><div class="para" label-module="para">著作</div></td><td align="middle" width="527"><div class="para" label-module="para">备注</div></td></tr><tr><td align="left" width="170"><div class="para" label-module="para">《实战经验录》</div></td><td align="left" width="527"><div class="para" label-module="para">与<a target="_blank" href="/item/%E7%BD%97%E5%BF%A0%E6%AF%85">罗忠毅</a>合编，1939年2月25日延安翻印了这本小册子。</div></td></tr><tr><td align="left" width="170" height="57"><div class="para" label-module="para">《激流归大海——回忆朱德同志和陈毅同志》</div></td><td align="left" width="527" height="57"><div class="para" label-module="para">上海人民出版社1979年版。</div></td></tr><tr><td align="left" width="170"><div class="para" label-module="para">《千万里转战》</div></td><td align="left" width="527"><div class="para" label-module="para">上海文艺出版社1987年版。</div></td></tr><tr><td align="left" width="170"><div class="para" label-module="para">《<a target="_blank" href="/item/%E7%B2%9F%E8%A3%95%E6%88%98%E4%BA%89%E5%9B%9E%E5%BF%86%E5%BD%95">粟裕战争回忆录</a>》</div></td><td align="left" width="527"><div class="para" label-module="para">记录整理：楚青。解放军出版社1988年11月第1版，1995年8月重印。2005年1月由知识产权出版社按需出版，增加了大环境下当时未能出版的 第二十章《粟裕谈淮海战役》（该文献20，000字）。</div></td></tr><tr><td align="left" width="170"><div class="para" label-module="para">《<a target="_blank" href="/item/%E7%B2%9F%E8%A3%95%E5%86%9B%E4%BA%8B%E6%96%87%E9%9B%86">粟裕军事文集</a>》</div></td><td align="left" width="527"><div class="para" label-module="para">主编：<a target="_blank" href="/item/%E5%AD%99%E5%85%8B%E9%AA%A5">孙克骥</a>，《粟裕军事文集》编辑组编，解放军出版社1989年7月第1版，1991年7月再版重印。</div></td></tr><tr><td align="left" width="170"><div class="para" label-module="para">《<a target="_blank" href="/item/%E7%B2%9F%E8%A3%95%E8%AE%BA%E8%8B%8F%E4%B8%AD%E6%8A%97%E6%88%98">粟裕论苏中抗战</a>》</div></td><td align="left" width="527"><div class="para" label-module="para">《粟裕军事文集》编辑组编，江苏人民出版社1993年第1版。</div></td></tr><tr><td align="left" width="170"><div class="para" label-module="para">《粟裕文选》</div></td><td align="left" width="527"><div class="para" label-module="para">（三卷本，181.2万字），《粟裕文选》编辑组编，军事科学出版社2004年9月版。<sup class="sup--normal" data-sup="12" data-ctrmap=":12,">
[12]</sup><a class="sup-anchor" name="ref_[12]_1827"> </a>
</div></td></tr></tbody></table><div class="anchor-list ">
<a name="3" class="lemma-anchor para-title"></a>
<a name="sub1827_3" class="lemma-anchor "></a>
<a name="人物成就" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>人物成就</h2>
</div>
<div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/1e71f7242b89dd2c4c088d1b?fr=lemma&amp;ct=single" target="_blank" title="淮海战役中的粟裕" style="width:220px;height:325.6px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/4b90f603738da977aa02c736b051f8198618e3fa?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="淮海战役中的粟裕" style="width:220px;height:325.6px;"/>
</a>
<span class="description">
淮海战役中的粟裕
</span>
</div>粟裕先后创建了浙南、<a target="_blank" href="/item/%E8%8B%8F%E5%8D%97/24746" data-lemmaid="24746">苏南</a>、<a target="_blank" href="/item/%E8%8B%8F%E4%B8%AD">苏中</a>、海上和苏 浙皖边根据地，而<a target="_blank" href="/item/%E8%8B%8F%E4%B8%AD%E6%8A%97%E6%97%A5%E6%A0%B9%E6%8D%AE%E5%9C%B0">苏中抗日根据地</a>的建立，使新四军在长江以北站稳了脚跟，改善了新四军在华中地区的战略态势，提高了根据地军民坚持敌后斗争的信心。<sup class="sup--normal" data-sup="13" data-ctrmap=":13,">
[13]</sup><a class="sup-anchor" name="ref_[13]_1827"> </a>
</div><div class="para" label-module="para">粟裕组建了苏浙公学并亲兼任校长，开办了各种短期训练队，培养适应现代战争要求的军政干部和参谋人才。</div><div class="para" label-module="para">粟裕组建的华中军区和华中野战军，实现了由游击兵团向正规军、由游击战向运动战的战略转变，在粟裕的指挥下，依托苏中解放区连续作战七次，并且“七战七捷”，鼓舞了解放区军民敢打必胜的信心，为解放战争初期的作战指导提供了实践经验。</div><div class="para" label-module="para">1945年12月25日，粟裕所指挥的高邮邵伯战役，创造了抗日战争期间一次战役歼灭日军人数的最高纪录（歼灭日军1100多人，歼灭伪军5000多人），打破了蒋军“开锁进门，长驱直入，直捣两淮”的企图，改善了华中南线战略态势，为后来的苏中战役创造了良好的战场条件。</div><div class="para" label-module="para">1948年11月6日，粟裕发起并指挥的<a target="_blank" href="/item/%E6%B7%AE%E6%B5%B7%E6%88%98%E5%BD%B9">淮海战役</a>，彻底粉碎了蒋介石在长江以北建立“重点防御”的计划，消灭了国民党在华东、中原战场上的精锐主力。<sup class="sup--normal" data-sup="4" data-ctrmap=":4,">
[4]</sup><a class="sup-anchor" name="ref_[4]_1827"> </a>
</div><div class="anchor-list ">
<a name="4" class="lemma-anchor para-title"></a>
<a name="sub1827_4" class="lemma-anchor "></a>
<a name="家庭成员" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>家庭成员</h2>
</div>
<table log-set-param="table_view" data-sort="sortDisabled"><tbody><tr><td align="middle" width="86"><div class="para" label-module="para">妻子</div></td><td align="left" width="611"><div class="para" label-module="para">楚青，原名詹永珠，出生于1923年3月，祖籍江苏省扬州市。1938年在皖南参加新四军。进入<a target="_blank" href="/item/%E6%96%B0%E5%9B%9B%E5%86%9B%E6%95%99%E5%AF%BC%E6%80%BB%E9%98%9F">新四军教导总队</a>第八队、新四军军部速记训练班学习，1939年3月就加入了中国共产党。1941年，18岁的楚青与34岁的粟裕在新四军司令部结为终身伴侣。<sup class="sup--normal" data-sup="14" data-ctrmap=":14,">
[14]</sup><a class="sup-anchor" name="ref_[14]_1827"> </a>
</div></td></tr><tr><td align="middle" width="86"><div class="para" label-module="para">岳父</div></td><td align="left" width="611"><div class="para" label-module="para">詹克明，是一位银行家。</div></td></tr><tr><td align="middle" width="86"><div class="para" label-module="para">长子</div></td><td align="left" width="611"><div class="para" label-module="para">粟戎生，中将军衔，1942年生于江苏扬州。1961年参加中国人民解放军，入哈尔滨军事工程学院导弹系学习。1966年加入中国共产党。原北京军区副司令员，1967年击落美侦察机作战中荣立三等功。</div></td></tr><tr><td align="middle" width="86"><div class="para" label-module="para">次子</div></td><td align="left" width="611"><div class="para" label-module="para">粟寒生，1947年出生。中国远洋公司副经理，曾任南方远洋总公司党委书记。</div></td></tr><tr><td align="middle" width="86"><div class="para" label-module="para">女儿</div></td><td align="left" width="611"><div class="para" label-module="para">粟惠宁，大校，原二炮研究院主任级。</div></td></tr><tr><td align="middle" width="86"><div class="para" label-module="para">女婿</div></td><td align="left" width="611"><div class="para" label-module="para">陈小鲁，<a target="_blank" href="/item/%E9%99%88%E6%AF%85">陈毅</a>元帅之子，1946年7月生于山东，文革前为<a target="_blank" href="/item/%E5%8C%97%E4%BA%AC%E7%AC%AC%E5%85%AB%E4%B8%AD%E5%AD%A6">北京第八中学</a>1966届高中毕业生。1969年 加入中国共产党，后任解放军第39军244团政治处主任。现为<a target="_blank" href="/item/%E5%8D%9A%E6%97%B6%E5%9F%BA%E9%87%91%E7%AE%A1%E7%90%86%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8">博时基金管理有限公司</a>、<a target="_blank" href="/item/%E6%B1%9F%E8%A5%BF%E9%95%BF%E8%BF%90%E8%82%A1%E4%BB%BD%E6%9C%89%E9%99%90%E5%85%AC%E5%8F%B8">江西长运股份有限公司</a>独立董事。</div></td></tr></tbody></table><div class="anchor-list ">
<a name="5" class="lemma-anchor para-title"></a>
<a name="sub1827_5" class="lemma-anchor "></a>
<a name="轶事典故" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>轶事典故</h2>
</div>
<div class="anchor-list ">
<a name="5_1" class="lemma-anchor para-title"></a>
<a name="sub1827_5_1" class="lemma-anchor "></a>
<a name="夫人赋诗" class="lemma-anchor "></a>
<a name="5-1" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>夫人赋诗</h3>
</div>
<div class="para" label-module="para">夫人<a target="_blank" href="/item/%E6%A5%9A%E9%9D%92">楚青</a>写诗一首，以寄托和粟裕共同战斗、生活四十多年的深情。</div><div class="para" label-module="para"><b>遣怀</b></div><div class="para" label-module="para">时晴时雨正清明，万里送君伴君行。</div><div class="para" label-module="para">宽慰似见忠魂笑，遣怀珍惜战友情。</div><div class="para" label-module="para">唯思跃马挥鞭日，但忆疆场捷报频。</div><div class="para" label-module="para">东南此刻花似锦，堪慰英灵一片心。<sup class="sup--normal" data-sup="15" data-ctrmap=":15,">
[15]</sup><a class="sup-anchor" name="ref_[15]_1827"> </a>
</div><div class="anchor-list ">
<a name="5_2" class="lemma-anchor para-title"></a>
<a name="sub1827_5_2" class="lemma-anchor "></a>
<a name="颅中弹片" class="lemma-anchor "></a>
<a name="5-2" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>颅中弹片</h3>
</div>
<div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/242dd42a2834349b5c9eae1fc9ea15ce37d3be94?fr=lemma&amp;ct=single" target="_blank" title="留在粟裕头颅中长 达五十四年之久的三块弹片" style="width:220px;height:145.07246376812px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/242dd42a2834349b5c9eae1fc9ea15ce37d3be94?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="留在粟裕头颅中长达五十四年之久的三块弹片" style="width:220px;height:145.07246376812px;"/>
</a>
<span class="description">
留在粟裕头颅中长达五十四年之久的三块弹片
</span>
</div>粟裕一生先后6次负伤。头部两次负伤，在武平战斗中，子弹从他右耳上侧头部颞骨穿过；在水南 作战中，被炮弹炸伤头部。手臂两次负伤，在硝石与敌作战中，他左臂负重伤留下残疾；在浙西遂安向皖赣边的转战中，他右臂中弹，新中国成立后才取出子弹。除此之外，1929年攻占宁都时，他臀部负伤；1936年在云合开展游击战中，他脚踝负伤。</div><div class="para" label-module="para">1984年2月5日他逝世后，家人从他火化的头颅骨灰中，竟发现了三块弹片。</div><div class="para" label-module="para">2003年，<a target="_blank" href="/item/%E5%86%9B%E4%BA%8B%E7%A7%91%E5%AD%A6%E9%99%A2">军事科学院</a>筹建院史馆，粟裕大将夫人楚青公开了这三块珍藏近20年的弹片。<sup class="sup--normal" data-sup="16" data-ctrmap=":16,">
[16]</sup><a class="sup-anchor" name="ref_[16]_1827"> </a>
</div><div class="anchor-list ">
<a name="5_3" class="lemma-anchor para-title"></a>
<a name="sub1827_5_3" class="lemma-anchor "></a>
<a name="授衔让帅" class="lemma-anchor "></a>
<a name="5-3" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>授衔让帅</h3>
</div>
<div class="para" label-module="para">1955年，授衔元帅、将军以“资历、威望、战功”为主要依据。毛泽东有意将粟裕封为元帅，但被粟裕推辞了，但毛泽东认为粟裕“大将”还是要当的，而且须为十大大将之首。粟裕对他的军衔问题看得很淡泊，并称：“评我大将，就是够高的了，要什么元帅呢？我只嫌高， 不嫌低。”<sup class="sup--normal" data-sup="17" data-ctrmap=":17,">
[17]</sup><a class="sup-anchor" name="ref_[17]_1827"> </a>
</div><div class="anchor-list ">
<a name="5_4" class="lemma-anchor para-title"></a>
<a name="sub1827_5_4" class="lemma-anchor "></a>
<a name="艰难平反" class="lemma-anchor "></a>
<a name="5-4" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>艰难平反</h3>
</div>
<div class="para" label-module="para"><a class="lemma-album layout-right nslog:10000206" title="粟裕大将历史照片" href="/pic/%E7%B2%9F%E8%A3%95/116084/439441/9e7ce6dc009667eccc116641?fr=lemma&amp;ct=cover" target="_blank" style="width:222px;" nslog-type="10000206">
<div class="album-wrap" style="width:220px;height:240.68px;">
<img class="picture" alt="粟裕大将历史照片" src="https://bkimg.cdn.bcebos.com/pic/810a19d8bc3eb135e26c5fc6a61ea8d3fd1f4420?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" style="width:220px;height:240.68px;"/>
</div>
<div class="description">
粟裕大将历史照片<span class="number">(20张)</span>
</div>
<div class="albumBg" style="width:220px;">
</div>
</a>1958年，粟裕因“反对反教条主义”一事在军委扩大会议中受到错误的批判。1978年，中共中央开始开始着手处理历史遗留下来的冤假错案，<a target="_blank" href="/item/%E9%82%93%E5%B0%8F%E5%B9%B3/116181" data-lemmaid="116181">邓小平</a>称“要了结1958年军委扩大会议这桩公案”。粟裕得知消息 去见了<a target="_blank" href="/item/%E5%8F%B6%E5%89%91%E8%8B%B1">叶剑英</a>副主席，表示要求平反。叶剑英称：“这件事应该解决一下，你写个报告给中央，我回京后同小平同志也说一下。”从叶剑英那得知邓小平也同意后，粟裕于1979年10月9日向中央正式写了申诉报告，要求撤销会议强加给他的一切 诬蔑不实之词，叶剑英在粟裕的报告上批示同意。但不知何因，此事却一直拖了下来。</div><div class="para" label-module="para">1983年，<a target="_blank" href="/item/%E8%83%A1%E8%80%80%E9%82%A6">胡耀邦</a>总书记对“粟裕的冤案”又进一步批示，中共中央、中央军委决定直接受理后，提出了具体的方案，并征求了粟裕本人的意见，决定即由中共中央指派代表同粟裕本人正式见面，但这一决定也未能付诸实施。粟裕于1984年故去。</div><div class="para" label-module="para">1994年12月25日，中央军委副主席<a target="_blank" href="/item/%E5%88%98%E5%8D%8E%E6%B8%85">刘华清</a>和<a target="_blank" href="/item/%E5%BC%A0%E9%9C%87/4596171" data-lemmaid="4596171">张震</a>联名发表了题为《追忆粟裕同志》的文章。文章同时在党中央机关报《<a target="_blank" href="/item/%E4%BA%BA%E6%B0%91%E6%97%A5%E6%8A%A5">人民日报</a>》和中央军委机关报《<a target="_blank" href="/item/%E8%A7%A3%E6%94%BE%E5%86%9B%E6%8A%A5">解放军报</a>》刊登。文章除了对粟裕的战绩和品德作了全面的评价外，特别明确指出：“1958年，粟裕同志在军委扩大会议上受到错误的批判，并因此长期受到不公正 的对待。这是历史上的一个失误。这个看法，也是中央军事委员会的意见。”<sup class="sup--normal" data-sup="18" data-ctrmap=":18,">
[18]</sup><a class="sup-anchor" name="ref_[18]_1827"> </a>
</div><div class="anchor-list ">
<a name="6" class="lemma-anchor para-title"></a>
<a name="sub1827_6" class="lemma-anchor "></a>
<a name="纪念场所" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>纪念场所</h2>
</div>
<div class="anchor-list ">
<a name="6_1" class="lemma-anchor para-title"></a>
<a name="sub1827_6_1" class="lemma-anchor "></a>
<a name="纪念堂" class="lemma-anchor "></a>
<a name="6-1" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>纪念堂</h3>
</div>
<div class="para" label-module="para"><a class="lemma-album layout-right nslog:10000206" title="粟裕" href="/pic/%E7%B2%9F%E8%A3%95/116084/16899264/0df431adcbef76099b6408dc2edda3cc7cd99e02?fr=lemma&amp;ct=cover" target="_blank" style="width:222px;" nslog-type="10000206">
<div class="album-wrap" style="width:220px;height:147.4px;">
<img class="picture" alt="粟裕" src="https://bkimg.cdn.bcebos.com/pic/0df431adcbef76099b6408dc2edda3cc7cd99e02?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" style="width:220px;height:147.4px;"/>
</div>
<div class="description">
粟裕<span class="number">(58张)</span>
</div>
<div class="albumBg" style="width:220px;">
</div>
</a></div><div class="para" label-module="para">“粟裕纪念堂”于1987年建于江苏省<a target="_blank" href="/item/%E4%B8%9C%E5%8F%B0%E5%B8%82">东台市</a><a target="_blank" href="/item/%E4%B8%89%E4%BB%93%E9%95%87">三仓镇</a>陵园北部，为古典苏式建筑。1995年12月，中央军委副主席<a target="_blank" href="/item/%E5%BC%A0%E9%9C%87/4596171" data-lemmaid="4596171">张震</a>同志为“ 粟裕纪念堂”提写了匾额。现纪念堂内陈列有粟裕半身胸像模型和其夫人楚青敬献的花篮，以及老同志的 题词多幅，<a target="_blank" href="/item/%E4%B8%9C%E5%8F%B0">东台</a>市委党史部门在这里举办 了《粟裕将军在三仓》和《东台英烈》的图片展览。1984年4月19日，粟裕夫人楚青偕子女来三仓烈士陵 园，将粟裕部分骨灰安葬于纪念塔的西南侧，并在墓穴周围栽植七株松柏，象征“<a target="_blank" href="/item/%E4%B8%83%E6%88%98%E4%B8%83%E6%8D%B7/5613229" data-lemmaid="5613229">七战七捷</a>”。当年6月，<a target="_blank" href="/item/%E4%B8%89%E4%BB%93%E7%83%88%E5%A3%AB%E9%99%B5%E5%9B%AD">三仓烈士陵园</a>在安葬处建起了水泥墓，并立碑纪念。<sup class="sup--normal" data-sup="15" data-ctrmap=":15,">
[15]</sup><a class="sup-anchor" name="ref_[15]_1827"> </a>
</div><div class="anchor-list ">
<a name="6_2" class="lemma-anchor para-title"></a>
<a name="sub1827_6_2" class="lemma-anchor "></a>
<a name="纪念馆" class="lemma-anchor "></a>
<a name="6-2" class="lemma-anchor "></a>
</div><div class="para-title level-3" label-module="para-title">
<h3 class="title-text"><span class="title-prefix">粟裕</span>纪念馆</h3>
</div>
<div class="para" label-module="para"><a class="lemma-album layout-right nslog:10000206" title="粟裕公园" href="/pic/%E7%B2%9F%E8%A3%95/116084/439815/ac4bd11373f08202eb2d9a9e4bfbfbedaa641bd7?fr=lemma&amp;ct=cover" target="_blank" style="width:222px;" nslog-type="10000206">
<div class="album-wrap" style="width:220px;height:293.33333333333px;">
<img class="picture" alt="粟裕公园" src="https://bkimg.cdn.bcebos.com/pic/ac4bd11373f08202eb2d9a9e4bfbfbedaa641bd7?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" style="width:220px;height:293.33333333333px;"/>
</div>
<div class="description">
粟裕公园<span class="number">(4张)</span>
</div>
<div class="albumBg" style="width:220px;">
</div>
</a></div><div class="para" label-module="para">“粟裕纪念馆”位于湖南省<a target="_blank" href="/item/%E4%BC%9A%E5%90%8C%E5%8E%BF">会同县</a>城南郊粟裕公园内，建成于1987年。1991年9月国 家主席<a target="_blank" href="/item/%E6%9D%A8%E5%B0%9A%E6%98%86">杨尚昆</a>亲笔题写馆名，属纪念型建筑。馆内陈列有粟裕铜像，还有江泽民等党和国家领导人为粟裕题的词以及一些珍贵的照片、文献等等。三仓烈士陵园在安葬处建起了水泥墓，并立碑纪念。</div><div class="para" label-module="para">公园小山顶立有粟裕纪念碑，碑名1986年由国防部长张爱萍题写，碑高11.2米，其中碑身7.6米， 象征粟裕享年76岁，碑身的四周还有<a target="_blank" href="/item/%E5%AE%8B%E4%BB%BB%E7%A9%B7">宋任穷</a>、<a target="_blank" href="/item/%E5%BC%A0%E9%9C%87/4596171" data-lemmaid="4596171">张震</a>、<a target="_blank" href="/item/%E6%AF%9B%E8%87%B4%E7%94%A8">毛致用</a>等人的题词。碑座中心安放着盛有粟裕部分骨灰的骨灰盒。粟裕纪念馆地理优越，交通通讯便捷。枝柳铁路穿城而过，209国道通过馆前。馆外青山绿水相映生辉，是参观旅游的重要景点。<sup class="sup--normal" data-sup="19" data-ctrmap=":19,">
[19]</sup><a class="sup-anchor" name="ref_[19]_1827"> </a>
</div><div class="anchor-list ">
<a name="7" class="lemma-anchor para-title"></a>
<a name="sub1827_7" class="lemma-anchor "></a>
<a name="人物评价" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>人物评价</h2>
</div>
<div class="para" label-module="para"><div class="lemma-picture text-pic layout-right" style="width:220px; float: right;">
<a class="image-link" nslog-type="9317" href="/pic/%E7%B2%9F%E8%A3%95/116084/0/d01373f082025aaf4ad4b64efeedab64024f1a55?fr=lemma&amp;ct=single" target="_blank" title="粟裕大将" style="width:220px;height:350.24539877301px;">
<img class="lazy-img" src="data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAMAAAAoyzS7AAAAGXRFWHRTb2Z0d2FyZQBBZG9iZSBJbWFnZVJlYWR5ccllPAAAAAZQTFRF9fX1AAAA0VQI3QAAAAxJREFUeNpiYAAIMAAAAgABT21Z4QAAAABJRU5ErkJggg==" data-src="https://bkimg.cdn.bcebos.com/pic/d01373f082025aaf4ad4b64efeedab64024f1a55?x-bce-process=image/resize,m_lfit,w_220,limit_1/format,f_auto" alt="粟裕大将" style="width:220px;height:350.24539877301px;"/>
</a>
<span class="description">
粟裕大将
</span>
</div>1946年2月，粟裕在组织大兵团作战中，他用兵灵活，不拘一格，被陈毅誉为“愈出愈奇，越打越妙”。<sup class="sup--normal" data-sup="20" data-ctrmap=":20,">
[20]</sup><a class="sup-anchor" name="ref_[20]_1827"> </a>
</div><div class="para" label-module="para">1946年8月28日，毛泽东发电报《华中野战军的作战经 验》：“粟裕指挥正确，既灵活，又勇敢，故能取得伟大胜利”。</div><div class="para" label-module="para">1949年，毛泽东说：“淮海战役，粟裕立了第一功”。</div><div class="para" label-module="para">1949年，刘伯承说：“粟裕同志智深勇沉，非常优秀，百战百胜，有古名将之风，是我军最优秀的 将领，是中国的战略家”。<sup class="sup--normal" data-sup="21" data-ctrmap=":21,">
[21]</sup><a class="sup-anchor" name="ref_[21]_1827"> </a>
</div><div class="para" label-module="para">1955年8月27日，粟裕被评为“<a target="_blank" href="/item/%E5%8D%81%E5%A4%A7%E5%A4%A7%E5%B0%86">十大大将</a>”之一。</div><div class="para" label-module="para">1988年10月，粟裕被中央军委评为“中国共产党36位开国军事家”之一。</div><div class="para" label-module="para">1992年解放军出版社王希先发表文章《浅谈粟裕的军事思想》说：“ 粟裕是杰出的军事家、战略家”。<sup class="sup--normal" data-sup="21" data-ctrmap=":21,">
[21]</sup><a class="sup-anchor" name="ref_[21]_1827"> </a>
</div><div class="anchor-list ">
<a name="8" class="lemma-anchor para-title"></a>
<a name="sub1827_8" class="lemma-anchor "></a>
<a name="影视形象" class="lemma-anchor "></a>
</div><div class="para-title level-2" label-module="para-title">
<h2 class="title-text"><span class="title-prefix">粟裕</span>影视形象</h2>
</div>
<table log-set-param="table_view" data-sort="sortDisabled"><tbody><tr><td align="middle" width="113"><div class="para" label-module="para">电视剧</div></td><td valign="top" align="left" width="584"><div class="para" label-module="para">《<a target="_blank" href="/item/%E7%B2%9F%E8%A3%95%E5%A4%A7%E5%B0%86">粟裕大将</a>》、《<a target="_blank" href="/item/%E4%B8%83%E6%88%98%E4%B8%83%E6%8D%B7">七战七捷</a>》、《<a target="_blank" href="/item/%E8%8B%B1%E9%9B%84%E5%AD%9F%E8%89%AF%E5%B4%AE">英雄孟良崮</a>》、《<a target="_blank" href="/item/%E8%B1%AB%E4%B8%9C%E4%B9%8B%E6%88%98">豫东之战</a>》、《<a target="_blank" href="/item/%E6%B5%8E%E5%8D%97%E6%88%98%E5%BD%B9">济南战役</a>》（系列电视剧《苏中保卫战》《<a target="_blank" href="/item/%E5%8D%8E%E4%B8%9C%E9%87%8E%E6%88%98%E5%86%9B">华东野战军</a>》）、《<a target="_blank" href="/item/%E9%87%8E%E6%88%98%E5%B8%88">野战师</a>》等。</div><div class="para" label-module="para"><b>还有在以下电视剧中出现过</b></div><div class="para" label-module="para">《<a target="_blank" href="/item/%E5%8F%A4%E5%9F%8E%E6%83%85%E6%81%A8">古城情恨</a>》、《<a target="_blank" href="/item/%E6%96%B0%E5%9B%9B%E5%86%9B">新四军</a>》、《<a target="_blank" href="/item/%E7%BA%A2%E6%97%A5">红日</a>》、《<a target="_blank" href="/item/%E8%A7%A3%E6%94%BE">解放</a>》、《<a target="_blank" href="/item/%E4%BA%95%E5%86%88%E5%B1%B1">井冈山</a>》、《<a target="_blank" href="/item/%E4%B8%8A%E5%B0%86%E8%AE%B8%E4%B8%96%E5%8F%8B">上将许世友</a>》、《<a target="_blank" href="/item/%E6%B5%B4%E8%A1%80%E5%9D%9A%E6%8C%81">浴血坚持</a>》、《<a target="_blank" href="/item/%E5%8F%B6%E6%8C%BA%E5%B0%86%E5%86%9B">叶挺将军</a>》、《<a target="_blank" href="/item/%E4%B8%9C%E6%96%B9">东方</a>》、《<a target="_blank" href="/item/%E5%86%B3%E6%88%98%E5%8D%97%E4%BA%AC">决战南京</a>》等，《<a target="_blank" href="/item/%E4%BA%AE%E5%89%91">亮剑</a>》中解放战争时期淮海战场上的华野代司令员正是粟裕。<sup class="sup--normal" data-sup="22" data-ctrmap=":22,">
[22]</sup><a class="sup-anchor" name="ref_[22]_1827"> </a>
</div></td></tr><tr><td align="middle" width="113"><div class="para" label-module="para">电 影</div></td><td valign="top" align="left" width="584"><div class="para" label-module="para">《<a target="_blank" href="/item/%E5%A4%A7%E5%86%B3%E6%88%98">大决战</a>（淮海战役）》、《 开国大典》、《<a target="_blank" href="/item/%E5%A4%A7%E8%BF%9B%E5%86%9B">大进军</a>（<a target="_blank" href="/item/%E5%A4%A7%E6%88%98%E5%AE%81%E6%B2%AA%E6%9D%AD">大战宁沪杭</a>）》等，《<a target="_blank" href="/item/%E9%BB%84%E6%A1%A5%E5%86%B3%E6%88%98">黄桥决战</a>》中的 谷盈原型也是粟裕。<sup class="sup--normal" data-sup="22" data-ctrmap=":22,">
[22]</sup><a class="sup-anchor" name="ref_[22]_1827"> </a>
</div></td></tr><tr><td align="middle" width="113"><div class="para" label-module="para">纪 录片</div></td><td valign="top" align="left" width="584"><div class="para" label-module="para">《<a target="_blank" href="/item/%E7%B2%9F%E8%A3%95%E5%A4%A7%E5%B0%86">粟裕大将</a>》、《<a target="_blank" href="/item/%E7%99%BE%E6%88%98%E7%BB%8F%E5%85%B8">百战经典</a>名将与名战 之常胜将军粟裕》、《中国记忆——大将粟裕》、《生死对决——淮海战役全纪录》、《百战经典雄师劲旅之坐断东南》、《<a target="_blank" href="/item/%E6%96%B0%E5%9B%9B%E5%86%9B/648515" data-lemmaid="648515">新四军</a>》等。<sup class="sup--normal" data-sup="22-24" data-ctrmap=":23,:22,:24,">
[22-24]</sup><a class="sup-anchor" name="ref_[22-24]_1827"> </a>
</div></td></tr></tbody></table><div class="anchor-list ">
<a name="album-list" class="lemma-anchor "></a>
</div><div class="album-list">
<div class="header">
<span class="title">词条图册</span>
<a class="more-link" href="/pic/%E7%B2%9F%E8%A3%95/116084?fr=lemma" target="_blank" nslog-type="10000204">更多图册<em></em></a>
</div>
<div class="scroller">
<div class="list">
</div>
</div>
<div class="footer">
</div>
</div>
<dl class="lemma-reference collapse nslog-area log-set-param" data-nslog-type="2" log-set-param="ext_reference">
<dt class="reference-title">参考资料</dt>
<dd class="reference-list-wrap">
<ul class="reference-list">
<li class="reference-item reference-item--type1 " id="reference-[1]-1827-wrap">
<span class="index">1.</span>
<a class="gotop anchor" name="refIndex_1_1827" id="refIndex_1_1827" title="向上跳转" href="#ref_[1]_1827">  </a>
<a rel="nofollow" href="/reference/116084/77e3E2I1IR80W176Qr2tZ0XFkxgZvThrExU6zEh7KCM_pUoIyoHbrxBpN47pJ1ldPvCuNnxHr3PChMDSEyVlGqzhxcjGfapSHuD_jI1AJCOOFe1ckocQsdiorsq_TaniPuca" target="_blank" class="text">十大大将：粟裕简介<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>[引用日期2020-03-10]</span></li><li class="reference-item reference-item--type1 " id="reference-[2]-1827-wrap">
<span class="index">2.</span>
<a class="gotop anchor" name="refIndex_2_1827" id="refIndex_2_1827" title="向上跳转" href="#ref_[2]_1827">  </a>
<a rel="nofollow" href="/reference/116084/7e226BJmhIr3xZ80VvVIacMXXTnOZb7EwSBd7DRLZZdYOJqG-TuVDEOrTdpRrMUtiLr1SiVm5m7VLXygBZFnX_ByBXlqCX6O1TOwsGPRf9husb4NPkYGWbpRGQ" target="_blank" class="text">粟裕生平及相关文章<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>[引用日期2013-10-16]</span></li><li class="reference-item reference-item--type1 " id="reference-[3]-1827-wrap">
<span class="index">3.</span>
<a class="gotop anchor" name="refIndex_3_1827" id="refIndex_3_1827" title="向上跳转" href="#ref_[3]_1827">  </a>
<a rel="nofollow" href="/reference/116084/1abd-eQkUZpWQDztEX78L9keYK1S9BrTNFfQzGaSUqyGvAUykAzEF4zxTPJXlfTBw2jvawC7lpDvBb45MlKGG_LGd4N-43fWqIWyDgPHpPuITFcV1A80yRM" target="_blank" class="text">中国人民解放军开国将帅名录<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>[引用日期2014-03-08]</span></li><li class="reference-item reference-item--type1 " id="reference-[4]-1827-wrap">
<span class="index">4.</span>
<a class="gotop anchor" name="refIndex_4_1827" id="refIndex_4_1827" title="向上跳转" href="#ref_[4]_1827">  </a>
<a rel="nofollow" href="/reference/116084/0b200Nz0MGs2DHaVAiJ-yojSiuON-6cEA9N0_XuH-RIHALL1n53ti4xGAOAcjcYUo8KMC7PZicbakt_yyLYMg6Xo5kQdgT7XK3yQgYTnPPEe6gCz3A" target="_blank" class="text">粟裕大事年表<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>．2006</span><span>[引用日期2013-07-07]</span></li><li class="reference-item reference-item--type1 " id="reference-[5]-1827-wrap">
<span class="index">5.</span>
<a class="gotop anchor" name="refIndex_5_1827" id="refIndex_5_1827" title="向上跳转" href="#ref_[5]_1827">  </a>
<a rel="nofollow" href="/reference/116084/d235Ip-1xZfTMBNvS9gIbyle8RLqgNyiWiB82hpE7zaIGACbaXrsJN7fylMGlang5BII14rBpdw-7-Xqu0WAm7MryVA8OG4_T62Hnab9umXOEZ-QG4R-" target="_blank" class="text">苏中反“扫荡”和基点争夺战*<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>[引用日期2013-10-16]</span></li><li class="reference-item reference-item--type1 " id="reference-[6]-1827-wrap">
<span class="index">6.</span>
<a class="gotop anchor" name="refIndex_6_1827" id="refIndex_6_1827" title="向上跳转" href="#ref_[6]_1827">  </a>
<a rel="nofollow" href="/reference/116084/8cd8mb0JKJWWYIwPUn1OAKomZ3mE6IgH3bEcx5YUCXbIoEQVwjr6esB5tKNZ0C4YuJ6UtZJkDAc0OCjZ1d_M-EUyZb2xSMqoA40KlQ8HQxKeGkgSRREFAZnK_9DvYg5w8JD7wTIpaIRdU8g" target="_blank" class="text">苏中战役<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>．2010-5-17</span><span>[引用日期2013-07-14]</span></li><li class="reference-item reference-item--type1 " id="reference-[7]-1827-wrap">
<span class="index">7.</span>
<a class="gotop anchor" name="refIndex_7_1827" id="refIndex_7_1827" title="向上跳转" href="#ref_[7]_1827">  </a>
<a rel="nofollow" href="/reference/116084/1a25k1lKivGLEdJhREIPK9qK_tfpN8xo0dX83suhE8O-y4m7Ytmn-KZY0tVNW5oDxBRbP5QnxFURqFBwNx4bYIQcjxcwkRq4VzczHjtnegNTVLK54i9jbPVSa2Ot5VAdPGIDqqk" target="_blank" class="text">豫东战役：粟裕如何终结蒋介石的中原梦图<span class="linkout"> </span></a>
<span class="site">．中华网</span><span>．2011-08-05</span><span>[引用日期2013-07-14]</span></li><li class="reference-item reference-item--type1 " id="reference-[8]-1827-wrap">
<span class="index">8.</span>
<a class="gotop anchor" name="refIndex_8_1827" id="refIndex_8_1827" title="向上跳转" href="#ref_[8]_1827">  </a>
<a rel="nofollow" href="/reference/116084/46b1759SAe6U6EvsbtDIZdTen2wyyyYxWQy633ZQQbpAH1ys_1XGrsm8v5aHZ3jqG-mz4AkfyashlzZTndG-YdqMIxrZ3zX3rtXY2c5lYXz_yuOzHn53jCQFomKJk6Fbamhu39oPZjkw8kUvhANrzCcwwxKoSA" target="_blank" class="text">淮海战役<span class="linkout"> </span></a>
<span class="site">．徐州网</span><span>．2010-10-10</span><span>[引用日期2014-06-22]</span></li><li class="reference-item reference-item--type1 " id="reference-[9]-1827-wrap">
<span class="index">9.</span>
<a class="gotop anchor" name="refIndex_9_1827" id="refIndex_9_1827" title="向上跳转" href="#ref_[9]_1827">  </a>
<a rel="nofollow" href="/reference/116084/b997HXWHZIBej9PfJR-WkAbFG61Ra_b1bt56lk1u-WOU-oq_ycZewY4J_HfXkYiLWcFY2WGaWhIqNWcjAThqIb6ltJyRjV6jFxAf5hwoFx5ug8hD6kzTo0K_u-8" target="_blank" class="text">群众路线:中国共产党“赶考”路上的生命线-<span class="linkout"> </span></a>
<span class="site">．中国干部学习网</span><span>．2010-01-01</span><span>[引用日期2014-06-17]</span></li><li class="reference-item reference-item--type1 " id="reference-[10]-1827-wrap">
<span class="index">10.</span>
<a class="gotop anchor" name="refIndex_10_1827" id="refIndex_10_1827" title="向上跳转" href="#ref_[10]_1827">  </a>
<a rel="nofollow" href="/reference/116084/7b77SdMbDCFk9oi1INvnJLWyZ5shD-ku-2W-9rIFmnRTc3C5nmXZ8BJEKYcrtFERRz4XVu0c6UWPigC8xjafDy_kDr20" target="_blank" class="text">战史今日：粟裕逝世 （2.5）<span class="linkout"> </span></a>
<span class="site">．腾讯网</span><span>．2010-02-05</span><span>[引用日期2012-12-08]</span></li><li class="reference-item reference-item--type1 more" id="reference-[11]-1827-wrap">
<span class="index">11.</span>
<a class="gotop anchor" name="refIndex_11_1827" id="refIndex_11_1827" title="向上跳转" href="#ref_[11]_1827">  </a>
<a rel="nofollow" href="/reference/116084/60d1vy331hSqN8RqkPhZynxPdSHWXn9cnSTwHfmv5kJPNsnX74x_cooZ1dfym1_a7DZadEHElzCD7dxtRoopFd-0H5vEciGG05VzOb8TiWsDyMSzQzA5_9g" target="_blank" class="text">粟裕<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>[引用日期2012-09-25]</span></li><li class="reference-item reference-item--type1 more" id="reference-[12]-1827-wrap">
<span class="index">12.</span>
<a class="gotop anchor" name="refIndex_12_1827" id="refIndex_12_1827" title="向上跳转" href="#ref_[12]_1827">  </a>
<a rel="nofollow" href="/reference/116084/0e2axIT4WbLm1nQ3mPp_Jw9mpgYjoLr30c3qM4kSZk4Bp7639qGBySKZv6X3MkaqIjhztWjV_j3sZQabFB-WptTFlsQS" target="_blank" class="text">粟裕大将百年诞辰纪 念活动9月举行<span class="linkout"> </span></a>
<span class="site">．腾讯网</span><span>．2007-08-14</span><span>[引用日期2012-12-08]</span></li><li class="reference-item reference-item--type1 more" id="reference-[13]-1827-wrap">
<span class="index">13.</span>
<a class="gotop anchor" name="refIndex_13_1827" id="refIndex_13_1827" title="向上跳转" href="#ref_[13]_1827">  </a>
<a rel="nofollow" href="/reference/116084/207fStzsGFGR7Zt-IbkmDVfA7karfFqV-KB2Pa3wfly1J8yA30QXtvHGdjdXvsvvYE6qECpbBreu-EU-xVctZH1BLC5WSi9dXamShKIpyu8dGW0kWCbH" target="_blank" class="text">苏中抗日根据地<span class="linkout"> </span></a>
<span class="site">．中国共产党新闻网</span><span>[引用日期2015-02-02]</span></li><li class="reference-item reference-item--type1 more" id="reference-[14]-1827-wrap">
<span class="index">14.</span>
<a class="gotop anchor" name="refIndex_14_1827" id="refIndex_14_1827" title="向上跳转" href="#ref_[14]_1827">  </a>
<a rel="nofollow" href="/reference/116084/d168_EvrDz6rvQ51Hkxd_4kLA28zyMqGlS95cAKMpeVF_RFENqODvSsheQBK6PKckB_W7YIkMeqOG18JgXH7SpBGF1SUV94y7V-M0-WsBFw" target="_blank" class="text">常胜将军决胜“情” 粟裕大将的爱情乐章<span class="linkout"> </span></a>
<span class="site">．中国共产党新闻网</span><span>[引用日期2015-01-26]</span></li><li class="reference-item reference-item--type1 more" id="reference-[15]-1827-wrap">
<span class="index">15.</span>
<a class="gotop anchor" name="refIndex_15_1827" id="refIndex_15_1827" title="向上跳转" href="#ref_[15]_1827">  </a>
<a rel="nofollow" href="/reference/116084/0716REJFwL4gFKa42agRlFRWh71-iiP2__6x_bcZhEWfC6n_BZdXYJlAR3QuALDDvlhId04hypWkE_BIFQgXu-8tLb5GVV_vANk" target="_blank" class="text">粟裕将军生平介绍<span class="linkout"> </span></a>
<span class="site">．族谱录纪念网</span><span>[引用日期2013-08-11]</span></li><li class="reference-item reference-item--type1 more" id="reference-[16]-1827-wrap">
<span class="index">16.</span>
<a class="gotop anchor" name="refIndex_16_1827" id="refIndex_16_1827" title="向上跳转" href="#ref_[16]_1827">  </a>
<a rel="nofollow" href="/reference/116084/3e30v1XEZ33uwb3o3gpLJ4bS7J5ncsrehfvuOWNYTHAowpTA8UfDSx2Oe10WHhtF36c3Gqtt_qGiuNtQ-gRp6L49y879CRObiYVVo5GHpD-jQjzOPpCl46g" target="_blank" class="text">粟裕大将头颅中的三块弹片（图）<span class="linkout"> </span></a>
<span class="site">．新华网</span><span>．2005-3-28</span><span>[引用日期2013-10-16]</span></li><li class="reference-item reference-item--type1 more" id="reference-[17]-1827-wrap">
<span class="index">17.</span>
<a class="gotop anchor" name="refIndex_17_1827" id="refIndex_17_1827" title="向上跳转" href="#ref_[17]_1827">  </a>
<a rel="nofollow" href="/reference/116084/81bfTINX21hNO2-2Lmf7GE7HlHkUWqSU7RtaC7gswR2A1eMvudiaN_tZ-6qJ5EpoywttrcYBqqx1Pw6AsEueKHvqiuOfhkKk99hX6h23HeQ7NMKdX74KAfE-0sE" target="_blank" class="text">“共和国第一大将”粟裕授衔的真相<span class="linkout"> </span></a>
<span class="site">．新华网</span><span>[引用日期2015-02-03]</span></li><li class="reference-item reference-item--type1 more" id="reference-[18]-1827-wrap">
<span class="index">18.</span>
<a class="gotop anchor" name="refIndex_18_1827" id="refIndex_18_1827" title="向上跳转" href="#ref_[18]_1827">  </a>
<a rel="nofollow" href="/reference/116084/4cb6KbENvyuYbFEkWJ_vtrr7ubOJ5A-lQqmtU654F-bsWgR7Kvnk_YVMuS7HZOTYzrIYjIfWoafBuz2JQlNKBE9EAQu7_F4" target="_blank" class="text">悲壮大将粟裕一生得不到平反的隐情<span class="linkout"> </span></a>
<span class="site">．红潮网</span><span>[引用日期2015-02-03]</span></li><li class="reference-item reference-item--type1 more" id="reference-[19]-1827-wrap">
<span class="index">19.</span>
<a class="gotop anchor" name="refIndex_19_1827" id="refIndex_19_1827" title="向上跳转" href="#ref_[19]_1827">  </a>
<a rel="nofollow" href="/reference/116084/11485lCZkrOkeHOCNtWefk9woz0qAXykIjMELRWFDaa-5vGzFxJyRMj8CBNC2KmSN7OisQw6k6-h49wsxwNg39NAu-9HO2Vr_7hMRFZMHnZsQmCHLA" target="_blank" class="text">粟裕同志纪念馆<span class="linkout"> </span></a>
<span class="site">．中国共产党新闻</span><span>[引用日期2013-08-11]</span></li><li class="reference-item reference-item--type1 more" id="reference-[20]-1827-wrap">
<span class="index">20.</span>
<a class="gotop anchor" name="refIndex_20_1827" id="refIndex_20_1827" title="向上跳转" href="#ref_[20]_1827">  </a>
<a rel="nofollow" href="/reference/116084/a0043vobuxSDCnS97HT-wnnmzCwOW9Spnwcg7E7HFTJ_jlHRfkDdgAWMKwpsi5GlxZeEhory_6V50gZw9opWjUureJ_55ELUkQqafK7mVD4y85mpJxhjHpwIMYBh3qdxQznUDM4BJv4" target="_blank" class="text">无产阶级革命家、军事家：粟裕大将[图]<span class="linkout"> </span></a>
<span class="site">．中华网</span><span>[引用日期2015-01-26]</span></li><li class="reference-item reference-item--type1 more" id="reference-[21]-1827-wrap">
<span class="index">21.</span>
<a class="gotop anchor" name="refIndex_21_1827" id="refIndex_21_1827" title="向上跳转" href="#ref_[21]_1827">  </a>
<a rel="nofollow" href="/reference/116084/cf13ykkK5Ngd2daOL74AypHJeF-uN6RPwuAQ6fiqNbe3raTgTutlfsSXgQKpIBH69bKenMDnMTOb6R_kSwdI6PTpAmHLEJfAt3DwMRm3hxuz13dIFdVoxcnJaS5aXEPqx0ZKmhwcQ7HK" target="_blank" class="text">特稿：请不要再用“大将模式”评价粟裕同志——粟裕同志一生中担任的党 政军职务概述（组图）<span class="linkout"> </span></a>
<span class="site">．中红网</span><span>[引用日期2015-01-26]</span></li><li class="reference-item reference-item--type1 more" id="reference-[22]-1827-wrap">
<span class="index">22.</span>
<a class="gotop anchor" name="refIndex_22_1827" id="refIndex_22_1827" title="向上跳转" href="#ref_[22]_1827">  </a>
<a rel="nofollow" href="/reference/116084/570b9_iVhtZcOkmn9tOAT8k34nwnruSKni1M8WkcEqlm4AqrQ953FAQbmUxpzIH3K0lUkcP3Km74AulCLIVaUG9Xfi0VbCqyWrc" target="_blank" class="text">光辉战绩<span class="linkout"> </span></a>
<span class="site">．族谱录纪念网</span><span>[引用日期2013-08-11]</span></li><li class="reference-item reference-item--type1 more" id="reference-[23]-1827-wrap">
<span class="index">23.</span>
<a class="gotop anchor" name="refIndex_23_1827" id="refIndex_23_1827" title="向上跳转" href="#ref_[23]_1827">  </a>
<a rel="nofollow" href="/reference/116084/0ff39jDdszBpplbDGW2YlfcB75N52m0e4OO_5HMO-CWTvakXegnMbpFO6kxhlrlGJ2ezZ3ZK9WiaXtkpelXupJA9d_Rx8ygB5w8arnpfb5_P7ZtAX7xS__5khXdvEi0" target="_blank" class="text">渡江战役<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>[引用日期2013-09-20]</span></li><li class="reference-item reference-item--type1 more" id="reference-[24]-1827-wrap">
<span class="index">24.</span>
<a class="gotop anchor" name="refIndex_24_1827" id="refIndex_24_1827" title="向上跳转" href="#ref_[24]_1827">  </a>
<a rel="nofollow" href="/reference/116084/ee69hMzJ8DvtqnKrkBWn-SoQD_JpStKuMRpl-HTH6sa2CmdI-JsnnPX8iZ4CYKsrFW-emTPSSXUNsOAA-ITK71mqTBO8dDRh_c38Mwc1f2Zji-FVwcmt" target="_blank" class="text">济南战役胜利的原因和重大意义<span class="linkout"> </span></a>
<span class="site">．人民网</span><span>．1948-10-04</span><span>[引用日期2013-09-20]</span></li></ul>
</dd>
<dd class="toggle">
<span class="text expand-text">展开全部</span>
<span class="text collapse-text">收起</span>
<em class="toggle-arrow"></em>
</dd>
</dl>
    `

	reader := strings.NewReader(html)
	doc, _ := goquery.NewDocumentFromReader(reader)
	rule := make(map[string]string)
	rule[`$t=ary;$e=div[class="lemma-summary"]>div[class="para"];`] = "summary"
	rule[`$t=map;$e=[class^="basicInfo-item"];$pk=dt[class="basicInfo-item name"];$pv=dd[class="basicInfo-item value"];`] = "profile"
	rule[`$t=map;$e=div[class="info"]>span;$pk=span[class="name"];$pv=span[class="title"];`] = "relation"
	rule[`$t=ary;$e=div[class="para"];`] = "para"
	rule[`$t=images;$e=a[class="image-link"];$pk=title;$pv=href;`] = "images"
	for k, v := range rule {
		fmt.Println("---------------------------")
		fmt.Println(v)
		fmt.Println("---------------------------")
		regType := regexp.MustCompile(`\$t\=(.*?);`)
		regElement := regexp.MustCompile(`\$e\=(.*?);`)
		rType := regType.FindStringSubmatch(k)
		rElement := regElement.FindStringSubmatch(k)
		if len(rType) < 2 {
			continue
		}
		if len(rElement) < 2 {
			continue
		}
		if "text" == rType[1] {
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				fmt.Println(s.Text())
			})
		} else if "ary" == rType[1] {
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				fmt.Println(trimSpace(s.Text()))
			})
		} else if "map" == rType[1] {
			regKeyClass := regexp.MustCompile(`\$pk\=\w*\[class\=\"(.*?)\"`)
			regValueClass := regexp.MustCompile(`\$pv\=\w*\[class\=\"(.*?)\"`)
			rKeyClass := regKeyClass.FindStringSubmatch(k)
			rValueClass := regValueClass.FindStringSubmatch(k)
			keyClass := ""
			valueClass := ""
			if len(rKeyClass) >= 2 {
				keyClass = rKeyClass[1]
			}
			if len(rValueClass) >= 2 {
				valueClass = rValueClass[1]
			}
			var siblingKey *goquery.Selection
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				if s.HasClass(keyClass) {
					siblingKey = s
				} else if s.HasClass(valueClass) {
					fmt.Println(trimSpace(siblingKey.Text()) + ": " + trimSpace(s.Text()))
				}
			})
		} else if "images" == rType[1] {
			regKey := regexp.MustCompile(`\$pk\=(.*?);`)
			regValue := regexp.MustCompile(`\$pv\=(.*?);`)
			rKey := regKey.FindStringSubmatch(k)
			rValue := regValue.FindStringSubmatch(k)
			key := ""
			value := ""
			if len(rKey) >= 2 {
				key = rKey[1]
			}
			if len(rValue) >= 2 {
				value = rValue[1]
			}
			doc.Find(rElement[1]).Each(func(i int, s *goquery.Selection) {
				title, exist := s.Attr(key)
				if exist {
					fmt.Println(trimSpace(title))
				}
                link, exist := s.Attr(value)
				if exist {
					fmt.Println(trimSpace(link))
				}
				c := colly.NewCollector(func(c *colly.Collector) {
					extensions.RandomUserAgent(c)
				},
				)

                c.OnHTML(`img[id="imgPicture"]`, func(e *colly.HTMLElement) {
					src := e.Attr("src")
                    fmt.Println(src)
				})

				c.OnError(func(r *colly.Response, e error) {
                    fmt.Println(e)
				})

                c.Visit("https://baike.baidu.com" + link)
			})
		}
	}
}
