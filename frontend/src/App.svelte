<script lang="ts">
  import svelteLogo from './assets/svelte.svg'
  import viteLogo from '/vite.svg'
    import Button from './lib/Button.svelte';
    import QuizCard from './lib/QuizCard.svelte';
    import { GameState, NetService, PacketTypes, type ChangeGameStatePacket, type PlayerJoinPacket, type TickPacket } from './service/net';
    import type { QuizQuestion } from './model/quiz';
    import type { Quiz } from './model/quiz';
    import type { Player } from './model/quiz';

  let quizzes: {_id: string, name: string}[] = [];

  let currentQuestion: QuizQuestion | null = null;
  let state = -1;
  let host = false;

  let tick = 0;

  let players: Player[] = [];

  let netService = new NetService();

  setTimeout(() => {
    netService.connect();
  }, 500);

  netService.onPacket((packet: any) => {
    console.log(packet);
    switch (packet.id) {
        case 2: {
            currentQuestion = packet.question;
            break;
        }

        case PacketTypes.ChangeGameState: {
            let data = packet as ChangeGameStatePacket;
            state = data.state;
            break;
        }

        case PacketTypes.PlayerJoin: {
            let data = packet as PlayerJoinPacket;
            players = [...players, data.player];
            break;
        }

        case PacketTypes.Tick: {
            let data = packet as TickPacket;
            tick = data.tick;
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

  function startGame() {
    netService.sendPacket({
        id: PacketTypes.StartGame
    })
  }

  function hostQuiz(quiz: Quiz) {
    host = true;
    netService.sendPacket({
        id: 1,
        quizId: quiz.id,
    })
  }
</script>

<Button on:click={getQuizzes}> Get quizzes </Button>
Message: {msg}

{#if state == -1}
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
{:else if state == GameState.Lobby}
    {#if host}
        <Button on:click={startGame}> Start Game </Button>
        <p> lobby state </p>
        {#each players as player}
            <p> {player.name} </p>
        {/each}
    {:else}
        <p> you have successfully connected </p>
    {/if}
{:else if state == GameState.Play}
    {#if host}
        Clock: {tick}
        {#if currentQuestion != null}
            <h2 class="text-4xl font-bold mt-8">{currentQuestion.name}</h2>
            <div class="flex"> 
                {#each currentQuestion.choices as choice}
                    <div class="flex-1 bg-blue-400 text-center font-bold text-2x1 text-white justify-center items-center p-8">
                        {choice.name}
                    </div>
                {/each}
            </div>
        {/if}
    {:else}
        press correct answer
    {/if}
{/if}