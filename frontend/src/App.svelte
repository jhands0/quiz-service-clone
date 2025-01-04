<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
    import Button from './lib/Button.svelte';
    import QuizCard from './lib/QuizCard.svelte';
    import { NetService } from './service/net';
    import type { QuizQuestion } from './model/quiz';
    import type { Quiz } from './model/quiz';

  let quizzes: {_id: string, name: string}[] = [];

  let currentQuestion: QuizQuestion | null = null;

  let netService = new NetService();
  netService.connect();
  netService.onPacket((packet: any) => {
    console.log(packet);
    switch (packet.id) {
        case 2: {
            currentQuestion = packet.question;
            break;
        }
    }
  });

  async function getQuizzes() {
    let response = await fetch("http://localhost:3000/api/quizzes")
    if (!response.ok) {
        alert("failed");
        return;
    }

    let json = await response.json();
    console.log(json);
  }

  let code = "";
  let name = "";
  let msg = "";

  function connect() {
    netService.sendPacket({
        id: 0,
        code: code,
        name: name,
    })
  }

  function hostQuiz(quiz: Quiz) {
    netService.sendPacket({
        id: 1,
        quizId: quiz.id,
    })
  }
</script>

<Button on:click={getQuizzes}> Get quizzes </Button>
Message: {msg}

<div>
{#each quizzes as quiz}
    <QuizCard on:host={() => hostQuiz(quiz)} quiz={quiz} />
{/each}
</div>

<input class="border" type="text" placeholder="Name" />
<input class="border" type="text" placeholder="Game code" />
<Button on:click={connect}> Join game </Button>

{#if currentQuestion != null}
    <h2 class="text-4x1 font-bold mt-8"> {currentQuestion.name} </h2>
    <div class="flex">
        {#each currentQuestion.choices as choice}
            <div class="flex-1 bg-blue-400 text-center font-bold text-2x1 text-white justify-center items-center p-8">
                {choice.name}
            </div>
        {/each}
    </div>
{/if}