<script lang="ts">
    import portrait from '../static/ME.jpg';
    import AboutMe from './AboutMe.svelte';
    import email from '../static/email.png';
    import phone from '../static/phone.png';

    export let name: string;
    export let title: string;
    export let aboutMe: string;
    export let personalMe: string;
    export let location: string;
</script>

<div class="header">
    <div class="left-side">
        <img class="picture" src={portrait} alt="me" />
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
                <span>{location}</span>
            </div>
        </div>
    </div>
    <div class="right-side">
        <AboutMe
            name={name}
            title={title}
            aboutMe={aboutMe}
            personalMe={personalMe}
        />
    </div>
</div>

<style lang="scss">
  .header {
    display: flex;
    flex-direction: row;
    gap: 2rem;
  }
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

  @media (max-width: 600px) {
    .header {
      flex-direction: column;
    }
  }
</style>
