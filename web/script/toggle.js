window.addEventListener('DOMContentLoaded', (event) => {
  me=document.querySelector("a.btn-cmt");
  if (!me) {return;}
  if (window.localStorage.getItem('comments-removed') === 'true') {
    document.getElementsByTagName("pre")[0].innerHTML = document.getElementsByTagName(
      "pre")[0].innerHTML.replaceAll(/<span class="c1">.*<\/span>/g, "")
      .replaceAll(/<span class="cm">.*<\/span>/g, "").replaceAll(/(\s*\n)/g, "\n")
  }
});

const doToggle = (evt) => {
  const ls = window.localStorage;
  const removed = () => ls.getItem("comments-removed")==="true";
  const setUnremoved = () => ls.setItem("comments-removed", "false");
  const setRemoved = () => ls.setItem("comments-removed", "true");
  const pre = () => document.getElementsByTagName("pre")[0];

  if(removed()) {
    setUnremoved();
    window.location.reload();
    return false;
  }

  setRemoved();

  pre().innerHTML = pre().innerHTML
    .replaceAll(/<span class="c1">.*<\/span>/g, "")
    .replaceAll(/<span class="cm">.*<\/span>/g, "")
    .replaceAll(/(\s*\n)/g, "\n");

  return false;
}