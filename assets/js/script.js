let datas = [];
// array digunakan untuk menyimpan data-data projects
function getData(event) {
  //  supaya ga nongol bentar
  event.preventDefault();
  //
  let projectName = document.getElementById('project-name').value;
  let startDate = document.getElementById('date-start').value;
  let endDate = document.getElementById('date-end').value;
  let description = document.getElementById('description').value;
  let nodeJs = document.getElementById('node-js').checked;
  let nextJs = document.getElementById('next-js').checked;
  let reactJs = document.getElementById('react-js').checked;
  let typeScript = document.getElementById('typescript').checked;
  let uploadImage = document.getElementById('upload-image').files;

  uploadImage = URL.createObjectURL(uploadImage[0]);

  let dataProject = {
    projectName,
    startDate,
    endDate,
    description,
    nodeJs,
    nextJs,
    reactJs,
    typeScript,
    uploadImage,
  };

  datas.push(dataProject);
  console.log(datas);
  showData();
}

const getDuration = (startDate, endDate) => {
  const duration = new Date(endDate) - new Date(startDate);

  let month = Math.floor(duration / (30 * 24 * 60 * 60 * 1000));
  if (month > 0) {
    return month + ' bulan';
  } else {
    let week = Math.floor(duration / (7 * 24 * 60 * 60 * 1000));
    if (week > 0) {
      return week + ' minggu';
    } else {
      let day = Math.floor(duration / (24 * 60 * 60 * 1000));
      if (day > 0) {
        return day + ' hari';
      } else {
        hour = 'kurang dari 1 hari';
      }
    }
  }
};
const showData = () => {
  //  buat hapus artikel yg udah ada di html sebelumnya
  document.getElementById('content-card').innerHTML = '';
  //
  for (let i = 0; i < datas.length; i++) {
    document.getElementById('content-card').innerHTML += `
    <div class="card">
      <div class="image-project">
        <img src="${datas[i].uploadImage}" />
      </div>
      <h3>${datas[i].projectName}</h3>
      <h5>durasi : ${getDuration(datas[i].startDate, datas[i].endDate)}</h5>
      <p>${datas[i].description}</p>
      <div class="tech-icon">
      ${datas[i].nodeJs ? '<img src="assets/icon/node.svg" />' : ''}
      ${datas[i].nextJs ? '<img src="assets/icon/next-js.svg" />' : ''}
      ${datas[i].reactJs ? '<img src="assets/icon/react.svg" />' : ''}
      ${datas[i].typeScript ? '<img src="assets/icon/typescript.svg" />' : ''}
      </div>
      <div class="manipulation">
        <a href="#">Edit</a>
        <a href="#">Delete</a>
      </div>
    </div>`;
  }
};
