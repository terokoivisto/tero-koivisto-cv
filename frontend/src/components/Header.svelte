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
                <span>{location}</span>
            </div>
        </div>
    </div>
    <div class="col-8">
        <AboutMe
            name={name}
            title={title}
            aboutMe={aboutMe}
            personalMe={personalMe}
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
