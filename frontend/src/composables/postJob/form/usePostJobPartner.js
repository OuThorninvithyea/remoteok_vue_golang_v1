import { ref, onMounted } from 'vue'
import axios from 'axios'

export function usePostJobPartner() {
    const partnerData = ref(null)
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const res = await axios.get('/api/postjob/partner')
            partnerData.value = res.data
        } catch (e) {
            error.value = e.message
        } finally {
            loading.value = false
        }
    })

    return { partnerData, loading, error }
}
