const testimonialData = [
  {
    author: 'Surya Elidanto',
    quote: 'Keren banget jasanya!',
    image: 'https://images.unsplash.com/photo-1570295999919-56ceb5ecca61?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Mnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60',
    rating: 5,
  },
  {
    author: 'Surya Elz',
    quote: 'Keren lah pokoknya!',
    image: 'https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60',
    rating: 4,
  },
  {
    author: 'Surya Gans',
    quote: 'The best pelayanannya!',
    image: 'https://images.unsplash.com/photo-1480429370139-e0132c086e2a?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Nnx8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60',
    rating: 4,
  },
  {
    author: 'Suryaaaa',
    quote: 'Oke lah!',
    image: 'https://media.istockphoto.com/id/1400280368/photo/happy-businessman-working-on-his-laptop-at-home-handsome-businessman-reading-an-email-on-his.jpg?b=1&s=170667a&w=0&k=20&c=mc9FiePkCPzKWRyexNf2lMo9BkDej_OpNloSDMNsutM=',
    rating: 3,
  },
  {
    author: 'Suryeah',
    quote: 'Apa apaan nih!',
    image: 'https://images.unsplash.com/photo-1615109398623-88346a601842?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTd8fG1hbnxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60',
    rating: 1,
  },
];

function allTestimonials() {
  let testimonialHTML = '';

  testimonialData.forEach((item) => {
    testimonialHTML += `<div class="testimonial">
                   <img src="${item.image}" class="profile-testimonial" />
                   <p class="quote">"${item.quote}"</p>
                   <p class="author">- ${item.author}</p>
                   <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
               </div>`;
  });

  document.getElementById('testimonials').innerHTML = testimonialHTML;
}

allTestimonials();

function filterTestimonials(rating) {
  let testimonialHTML = '';

  // rating : 1

  const testimonialFiltered = testimonialData.filter((item) => {
    return item.rating === rating;
  });

  // [
  //     {
  //         author: "Kevin Pratama",
  //         quote: "Apasih!",
  //         image: "https://images.unsplash.com/photo-1568602471122-7832951cc4c5?ixlib=rb-4.0.3&ixid=MnwxMjA3fDB8MHxzZWFyY2h8M3x8bWFufGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60",
  //         rating: 1
  //     },
  // ]

  if (testimonialFiltered.length === 0) {
    testimonialHTML = `<h1> Data not found! </h1>`;
  } else {
    testimonialFiltered.forEach((item) => {
      testimonialHTML += `<div class="testimonial">
              <img src="${item.image}" class="profile-testimonial" />
              <p class="quote">"${item.quote}"</p>
              <p class="author">- ${item.author}</p>
              <p class="author">${item.rating} <i class="fa-solid fa-star"></i></p>
          </div>`;
    });
  }

  document.getElementById('testimonials').innerHTML = testimonialHTML;
}
