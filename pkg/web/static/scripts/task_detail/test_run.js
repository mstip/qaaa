import { MethodEl, TestRunEl, TestRunResultEl, UrlEl } from "./elements.js";

export default function TestRun() {
  TestRunEl.onclick = function (event) {
    event.preventDefault();
    const method = MethodEl.value;
    const url = UrlEl.value;
    TestRunResultEl.innerText = "request send, waiting for response ....";

    axios
      .post("/api/task/testrun", { method, url })
      .then(function (response) {
        TestRunResultEl.innerText = `============================== Status ==============================\n`;
        TestRunResultEl.innerText += JSON.stringify(
          {
            status: response.status,
            statusTest: response.statusText,
            headers: response.headers,
          },
          0,
          2
        );
        TestRunResultEl.innerText += `\n============================== Data ==============================\n`;
        TestRunResultEl.innerText += JSON.stringify(response.data, 0, 2);
      })
      .catch(function (error) {
        TestRunResultEl.innerText = JSON.stringify(error, 0, 2);
      });
  };
}
