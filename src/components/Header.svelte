<script>
    import portrait from '../static/ME.jpg';
    import AboutMe from './AboutMe.svelte';
    import email from '../static/email.png';
    import phone from '../static/phone.png';
</script>

<div>
    <div class="col-4">
        <img src={portrait} alt="me" />
        <div class="contacts">
            <div class="contact">
                <iconify-icon icon="ic:baseline-email" />
                <img src={email} alt="email" />
            </div>
            <div class="contact">
                <iconify-icon icon="ic:baseline-phone" />
                <img src={phone} alt="phone" />
            </div>
            <div class="contact">
                <iconify-icon icon="mdi:location" />
                <span>Espoo, Finland</span>
            </div>
        </div>
    </div>
    <div class="col-8">
        <AboutMe
            name="Tero Koivisto"
            title="Senior Fullstack Developer"
            aboutMe="I'm a 28-year-old Senior Fullstack Developer who excels at adapting to challenges of all sorts and scales by breaking down larger problems into smaller easily comprehensible sub-problems, to deliver the solution that satisfies the requirements. I'm a quick learner and I'm always ready to expand beyond my current knowledge and skillset, if need be. Beyond writing code I always try to understand the underlying reason why we want to do something, and combine the business requirements to the current technical solution in a robust manner. I'm very used to speaking English with colleagues."
            personalMe="Outside my worklife, I spend time listening to music, playing various games (FPS and survival are my favorites), writing code for GTA V roleplaying server and during the summer I try my best to play some golf. I'm also a car enthusiast (yes, the ones using ancient technology that go brrrr), a semi-hardcore coffee lover and I like cooking all kinds of foods."
        />
    </div>
</div>

<style lang="scss">
  .contacts {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;

    .contact {
      display: flex;
      flex-direction: row;
      gap: 0.5rem;
      align-items: center;
    }
  }
</style>
