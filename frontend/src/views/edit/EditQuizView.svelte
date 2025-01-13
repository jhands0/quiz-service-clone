<script lang="ts">
    import type { Quiz } from "../../model/quiz";
    import { apiService } from "../../service/api";

    export let params: Record<string, string>

    let quiz: Quiz | null;

    (async function () {
        quiz = await apiService.getQuizById(params["quizId"]);
    })();
</script>