import { ref, onMounted } from 'vue'
import axios from 'axios'

export function usePostJobJobDetails() {
    const jobDetailsData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/job_details')
            jobDetailsData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { jobDetailsData, loading, error }
}
