import { ref, onMounted } from 'vue'
import axios from 'axios'

export function usePostJobHeader() {
    const headerData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/header')
            headerData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { headerData, loading, error }
}
